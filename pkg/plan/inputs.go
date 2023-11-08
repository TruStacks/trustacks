package plan

type InputField string

type InputFieldSchema struct {
	Type        string      `json:"type"`
	Pattern     string      `json:"pattern,omitempty"`
	Description string      `json:"description"`
	Default     interface{} `json:"default,omitempty"`
}

type input interface {
	Schema() InputFieldSchema
}

const ContainerRegistry InputField = "CONTAINER_REGISTRY"

type ContainerRegistryInput struct{}

func (input ContainerRegistryInput) Schema() InputFieldSchema {
	return InputFieldSchema{
		Type:        "String",
		Description: "The fully qualified URL of the container registry that will be used for both authentiation and as the destination where application container images will be be pushed.",
	}
}

const ContainerRegistryUsername InputField = "CONTAINER_REGISTRY_USERNAME"

type ContainerRegistryUsernameInput struct{}

func (input ContainerRegistryUsernameInput) Schema() InputFieldSchema {
	return InputFieldSchema{
		Type:        "String",
		Description: "The container registry username.",
	}
}

const ContainerRegistryPassword InputField = "CONTAINER_REGISTRY_PASSWORD"

type ContainerRegistryPasswordInput struct{}

func (input ContainerRegistryPasswordInput) Schema() InputFieldSchema {
	return InputFieldSchema{
		Type:        "String",
		Description: "The container registry password.",
	}
}

const SonarqubeToken InputField = "SONARQUBE_TOKEN"

type SonarqubeTokenInput struct{}

func (input SonarqubeTokenInput) Schema() InputFieldSchema {
	return InputFieldSchema{
		Type:        "String",
		Description: "The sonarqube access token.",
	}
}

const ArgoCDServer InputField = "ARGOCD_SERVER"

type ArgoCDServerInput struct{}

func (input ArgoCDServerInput) Schema() InputFieldSchema {
	return InputFieldSchema{
		Type:        "String",
		Description: "The ArgoCD server URL",
	}
}

const ArgoCDAuthToken InputField = "ARGOCD_AUTH_TOKEN"

type ArgoCDAuthTokenInput struct{}

func (input ArgoCDAuthTokenInput) Schema() InputFieldSchema {
	return InputFieldSchema{
		Type:        "String",
		Description: "The ArgoCD authentication token",
	}
}

const GithubToken InputField = "GITHUB_TOKEN"

type GithubTokenInput struct{}

func (input GithubTokenInput) Schema() InputFieldSchema {
	return InputFieldSchema{
		Type:        "String",
		Description: "The ArgoCD authentication token",
	}
}

var vars = map[string]input{
	"CONTAINER_REGISTRY":          ContainerRegistryInput{},
	"CONTAINER_REGISTRY_USERNAME": ContainerRegistryUsernameInput{},
	"CONTAINER_REGISTRY_PASSWORD": ContainerRegistryPasswordInput{},
	"SONARQUBE_TOKEN":             SonarqubeTokenInput{},
	"ARGOCD_SERVER":               ArgoCDServerInput{},
	"ARGOCD_AUTH_TOKEN":           ArgoCDAuthTokenInput{},
	"GITHUB_TOKEN":                GithubTokenInput{},
}

func GetInput(name string) input {
	return vars[name]
}
