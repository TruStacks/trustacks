package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"filippo.io/age"
	firebase "firebase.google.com/go"
	"github.com/charmbracelet/log"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/trustacks/pkg/actions"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
	"go.mozilla.org/sops/v3/decrypt"
	"golang.org/x/crypto/bcrypt"
	cryptossh "golang.org/x/crypto/ssh"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	projectID     = os.Getenv("FIREBASE_PROJECT_ID")
	jwtSigningKey = []byte(os.Getenv("JWT_SIGNING_KEY"))
)

type rpcHandler struct {
	projectID string
}

func (rpc *rpcHandler) newFrebaseClient() (*firestore.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}
	return app.Firestore(ctx)
}

func (rpc *rpcHandler) cloneSource(url, privateKey, path string) error {
	cloneOptions := &git.CloneOptions{
		URL:          url,
		SingleBranch: true,
		Depth:        0,
		Progress:     os.Stdout,
		Tags:         git.AllTags,
	}
	if privateKey != "" {
		publicKeys, err := ssh.NewPublicKeys("git", []byte(privateKey), "")
		if err != nil {
			return err
		}
		publicKeys.HostKeyCallback = cryptossh.InsecureIgnoreHostKey()
		cloneOptions.Auth = publicKeys
	}
	if _, err := git.PlainClone(path, false, cloneOptions); err != nil {
		return err
	}
	return nil
}

func (rpc *rpcHandler) userExists(username string) (bool, error) {
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return false, err
	}
	defer client.Close()
	doc, err := client.Collection("users").Doc(username).Get(context.Background())
	if err != nil && status.Code(err) != codes.NotFound {
		return false, nil
	}
	if doc.Exists() {
		return true, nil
	}
	return false, nil
}

func (rpc *rpcHandler) NewUser(username, password, registrationKey string) error {
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return err
	}
	defer client.Close()
	userExists, err := rpc.userExists(username)
	if err != nil {
		return err
	}
	if userExists {
		return errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = client.Collection("users").Doc(username).Set(context.Background(), map[string]interface{}{
		"password": hashedPassword,
	})
	return err
}

func (rpc *rpcHandler) validateSessionToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSigningKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("session token is invalid")
	}
	return nil
}

func (rpc *rpcHandler) getSessionTokenUser(tokenString string) (string, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSigningKey, nil
	})
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}

func (rpc *rpcHandler) RefreshSessionToken(tokenString string) (string, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("session token is invalid")
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		Issuer:    "trustacks",
		Subject:   claims.Subject,
	})
	return newToken.SignedString(jwtSigningKey)
}

func (rpc *rpcHandler) NewSessionToken(username, password string) (string, error) {
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return "", err
	}
	defer client.Close()
	doc, err := client.Collection("users").Doc(username).Get(context.Background())
	if err != nil && status.Code(err) != codes.NotFound {
		return "", err
	}
	if !doc.Exists() {
		return "", errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword(doc.Data()["password"].([]byte), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
		Issuer:    "trustacks",
		Subject:   username,
	})
	return token.SignedString(jwtSigningKey)
}

func (rpc *rpcHandler) NewActionPlan(url, path, privateKey, token string) error {
	if err := rpc.validateSessionToken(token); err != nil {
		return err
	}
	user, err := rpc.getSessionTokenUser(token)
	if err != nil {
		return err
	}
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return err
	}
	defer client.Close()
	userExists, err := rpc.userExists(user)
	if err != nil {
		return err
	}
	if !userExists {
		return errors.New("user does not exist")
	}
	re, err := regexp.Compile(`([^/]+$)`)
	if err != nil {
		return err
	}
	source, err := os.MkdirTemp("", "action-plan-source-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(source)
	matches := re.FindAllString(url, -1)
	if len(matches) < 1 {
		return errors.New("unable to derive name from git url")
	}
	name := strings.ToLower(strings.ReplaceAll(matches[0], ".git", ""))
	if path != "" {
		name = fmt.Sprintf("%s-%s", name, strings.ToLower(strings.ReplaceAll(path, "/", "")))
	}
	if err := rpc.cloneSource(url, privateKey, source); err != nil {
		return err
	}
	spec, err := engine.New().CreateActionPlan(filepath.Join(source, path), false)
	if err != nil {
		return err
	}
	plan := map[string]interface{}{}
	if err := json.Unmarshal([]byte(spec), &plan); err != nil {
		return err
	}
	data := map[string]interface{}{
		"name":       name,
		"user":       user,
		"plan":       plan,
		"repository": url,
	}
	if path != "" {
		data["path"] = path
	}
	_, _, err = client.Collection("actionplans").Add(context.Background(), data)
	return err
}

func (rpc *rpcHandler) ListActionPlans(token string) ([]map[string]interface{}, error) {
	if err := rpc.validateSessionToken(token); err != nil {
		return nil, err
	}
	user, err := rpc.getSessionTokenUser(token)
	if err != nil {
		return nil, err
	}
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	userExists, err := rpc.userExists(user)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, errors.New("user does not exist")
	}
	actionPlans := []map[string]interface{}{}
	iter := client.Collection("actionplans").Where("user", "==", user).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		data := doc.Data()["plan"].(map[string]interface{})
		actionPlan := map[string]interface{}{"name": doc.Data()["name"], "repository": doc.Data()["repository"]}
		if _, ok := data["actions"]; ok {
			actionPlan["actions"] = data["actions"]
		}
		if _, ok := data["fields"]; ok {
			actionPlan["fields"] = data["fields"]
		}
		if _, ok := data["exclusions"]; ok {
			actionPlan["exclusions"] = data["exclusions"]
		}
		actionPlans = append(actionPlans, actionPlan)
	}
	return actionPlans, nil
}

type GetActionPlanSpec struct {
	Path   string                 `json:"path"`
	Plan   []*plan.ActionSpec     `json:"spec"`
	Inputs map[string]interface{} `json:"inputs"`
}

func (rpc *rpcHandler) GetActionPlan(name, token string) (map[string]interface{}, error) {
	if err := rpc.validateSessionToken(token); err != nil {
		return nil, err
	}
	user, err := rpc.getSessionTokenUser(token)
	if err != nil {
		return nil, err
	}
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	userExists, err := rpc.userExists(user)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, errors.New("user does not exist")
	}
	iter := client.Collection("actionplans").
		Where("user", "==", user).
		Where("name", "==", name).
		Limit(1).
		Documents(context.Background())
	doc, err := iter.Next()
	if err != nil {
		return nil, err
	}
	if doc == nil {
		return nil, errors.New("action plan does not exist")
	}
	return doc.Data(), nil
}

func (rpc *rpcHandler) DeleteActionPlan(name string, token string) error {
	if err := rpc.validateSessionToken(token); err != nil {
		return err
	}
	user, err := rpc.getSessionTokenUser(token)
	if err != nil {
		return err
	}
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return err
	}
	defer client.Close()
	userExists, err := rpc.userExists(user)
	if err != nil {
		return err
	}
	if !userExists {
		return errors.New("user does not exist")
	}
	_, err = client.Collection("actionplans").Doc(name).Delete(context.Background())
	return err
}

func (rpc *rpcHandler) SetExcludedActions(name string, exclusions []string, token string) error {
	if err := rpc.validateSessionToken(token); err != nil {
		return err
	}
	user, err := rpc.getSessionTokenUser(token)
	if err != nil {
		return err
	}
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return err
	}
	defer client.Close()
	userExists, err := rpc.userExists(user)
	if err != nil {
		return err
	}
	if !userExists {
		return errors.New("user does not exist")
	}
	iter := client.Collection("actionplans").
		Where("user", "==", user).
		Where("name", "==", name).
		Limit(1).
		Documents(context.Background())
	doc, err := iter.Next()
	if err != nil {
		return err
	}
	if doc == nil {
		return errors.New("action plan does not exist")
	}
	data := doc.Data()
	data["exclusions"] = exclusions
	_, err = client.Collection("actionplans").Doc(doc.Ref.ID).Set(context.Background(), data)
	return err
}

func (rpc *rpcHandler) SaveStackContext(name, ageSecretKey string, ctx map[string]interface{}, token string) error {
	if err := rpc.validateSessionToken(token); err != nil {
		return err
	}
	user, err := rpc.getSessionTokenUser(token)
	if err != nil {
		return err
	}
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return err
	}
	defer client.Close()
	userExists, err := rpc.userExists(user)
	if err != nil {
		return err
	}
	if !userExists {
		return errors.New("user does not exist")
	}
	f, err := os.CreateTemp("", "stack-context-")
	if err != nil {
		return err
	}
	ctxJson, err := json.Marshal(ctx)
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if _, err := f.Write([]byte(ctxJson)); err != nil {
		return err
	}
	f.Seek(0, 0)
	b := make([]byte, 1000)
	if _, err := f.Read(b); err != nil {
		return err
	}
	identity, err := age.ParseX25519Identity(strings.ReplaceAll(ageSecretKey, "\n", ""))
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer([]byte(""))
	cmd := exec.Command("sops", "-e", "--age", identity.Recipient().String(), "--input-type", "json", f.Name())
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return err
	}
	data := map[string]interface{}{}
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		return err
	}
	_, err = client.Collection("stacks").Doc(name).Set(context.Background(), map[string]interface{}{
		"name": name,
		"user": user,
		"data": data,
	})
	return err
}

func (rpc *rpcHandler) ListStackContexts(ageSecretKey string, decryptInputs bool, token string) (map[string]interface{}, error) {
	if err := rpc.validateSessionToken(token); err != nil {
		return nil, err
	}
	user, err := rpc.getSessionTokenUser(token)
	if err != nil {
		return nil, err
	}
	stackInputs := map[string]interface{}{}
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return stackInputs, err
	}
	defer client.Close()
	if decryptInputs {
		_, err = age.ParseX25519Identity(strings.ReplaceAll(ageSecretKey, "\n", ""))
		if err != nil {
			return stackInputs, errors.New("age secret key could not be parsed")
		}
	}
	defer os.Unsetenv("SOPS_AGE_KEY")
	iter := client.Collection("stacks").Where("user", "==", user).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		data := doc.Data()["data"].(map[string]interface{})
		if decryptInputs && ageSecretKey != "" {
			dataJson, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			os.Setenv("SOPS_AGE_KEY", strings.ReplaceAll(ageSecretKey, "\n", ""))
			encryptedDataJson, err := decrypt.Data(dataJson, "json")
			if err != nil {
				return nil, fmt.Errorf("failed decrypting stack context inputs: confirm that you have the correct secret key")
			}
			if err := json.Unmarshal(encryptedDataJson, &data); err != nil {
				return nil, err
			}
		}
		delete(data, "sops")
		stackInputs[doc.Ref.ID] = data
	}
	return stackInputs, nil
}

func (rpc *rpcHandler) GetStackContext(name, ageSecretKey, token string) (map[string]interface{}, error) {
	if err := rpc.validateSessionToken(token); err != nil {
		return nil, err
	}
	user, err := rpc.getSessionTokenUser(token)
	if err != nil {
		return nil, err
	}
	client, err := rpc.newFrebaseClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	iter := client.Collection("stacks").
		Where("user", "==", user).
		Where("name", "==", name).
		Limit(1).
		Documents(context.Background())
	doc, err := iter.Next()
	if err != nil {
		return nil, err
	}
	data := doc.Data()["data"].(map[string]interface{})
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	os.Setenv("SOPS_AGE_KEY", strings.ReplaceAll(ageSecretKey, "\n", ""))
	defer os.Unsetenv("SOPS_AGE_KEY")
	encryptedDataJson, err := decrypt.Data(dataJson, "json")
	if err != nil {
		return nil, fmt.Errorf("failed decrypting stack context inputs: confirm that you have the correct secret key")
	}
	if err := json.Unmarshal(encryptedDataJson, &data); err != nil {
		return nil, err
	}
	delete(data, "sops")
	return data, nil
}

func StartServer(host, port string) error {
	rpcServer := jsonrpc.NewServer()
	rpcServer.Register("v1", &rpcHandler{})
	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		rpcServer.ServeHTTP(w, r)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK\n"))
	})
	log.Info(fmt.Sprintf("starting server on %s:%s", host, port))
	return http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}
