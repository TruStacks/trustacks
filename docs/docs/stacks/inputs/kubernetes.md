---
title: Kubernetes
slug: /inputs/kubernetes
---

# Kubernetes Inputs

Kubernetes inputs provide parameters for interacting with [kubernetes](https://kubernetes.io/) clusters.

## Fields

| name | type | description |
| - | - |-|
| KubernetesStagingKubeconfig | string | the kubeconfig file contents for kubernetes cluster used for staging the application. *See [above](#generating-a-kubeconfig) to create a compliant kubeconfig* |
| KubernetesNamespace | string | the namespace in the kubernetes cluster where the application will be deployed. |

## Generating a kubeconfig

TruStacks requires a [client certifcate](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#x509-client-certs) based kubernetes configuration. Use the following script to create one:

:::info
The following script requires an existing kubernetes cluster and the `kubectl` command line utility.
:::

```bash
#!/bin/bash

context="$1"
user="$2"

# create an rsa key and certificate signing request.
openssl genrsa -out $user.key 2048
openssl req -new -key $user.key -out $user.csr -subj="/CN=$user"

# create a certificate signing request resource in the kubernetes cluster.
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: $user
spec:
  request: $(cat $user.csr | base64 | tr -d "\n")
  signerName: kubernetes.io/kube-apiserver-client
  expirationSeconds: 86400
  usages:
  - client auth
EOF

# approve the csr and save the resulting certificate.
kubectl certificate approve $user
kubectl get csr $user -o jsonpath='{.status.certificate}'| base64 -d > $user.crt

# create the user cluster role binding.
kubectl create clusterrolebinding $user-cluster-admin --clusterrole=cluster-admin --user=$user --dry-run=client -o yaml | kubectl apply -f -

# collect the cluster details from the desired cluster context.
cluster_template="{{range \$context := .contexts}}{{if eq .name \"$context\"}}{{.context.cluster}}{{end}}{{end}}"
cluster=$(kubectl config view --flatten --template="$cluster_template")
cluster_server_template="{{range \$cluster := .clusters}}{{if eq .name \"$cluster\"}}{{.cluster.server}}{{end}}{{end}}"
cluster_server=$(kubectl config view --flatten --template="$cluster_server_template")
cluster_certificate_authority_data_template="{{range \$cluster := .clusters}}{{if eq .name \"$cluster\"}}{{.cluster.server}}{{end}}{{end}}"
cluster_certificate_authority_data=$(kubectl config view --flatten --template="$cluster_certificate_authority_data_template")

# create the kubeconfig.
export KUBECONFIG=/tmp/trustacks-$user-kubeconfig
touch $KUBECONFIG
kubectl config set-cluster $cluster --server=$cluster_server
kubectl config set-credentials $user --client-key=$user.key --client-certificate=$user.crt --embed-certs=true
kubectl config set-context $context --cluster=$cluster --user=$user
kubectl config use-context $context
kubectl config view --flatten > $context-$user-kubeconfig.yaml
kubectl config view --flatten -o json > $context-$user-kubeconfig.json
echo "KUBECONFIG created at $context-$user-kubeconfig.yaml(.json)"

# clean up script artifacts.
rm $KUBECONFIG $user.key $user.csr $user.crt
```

:::tip
The `context`(first argument) must come be an existing context in your kubeconfig. 

Run the following command to view your existing contexts: 
```
kubectl config get-contexts
``` 

The `user`(second argument) argument is the user that will be added to kubeconfig. The user requires no pre-existing service account or user in the kubernetes cluster.
:::
