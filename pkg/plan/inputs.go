package plan

type InputField string

type input interface {
	Description() string
	Link() string
}

const ContainerRegistry InputField = "CONTAINER_REGISTRY"

type ContainerRegistrySpec struct{}

func (input ContainerRegistrySpec) Description() string {
	return `
  The fully qualified URL of the container registry that will be used for both authentiation and as the destination 
  where application container images will be be pushed.
`
}

func (input ContainerRegistrySpec) Link() string {
	return "container"
}

const ContainerRegistryUsername InputField = "CONTAINER_REGISTRY_USERNAME"

type ContainerRegistryUsernameSpec struct{}

func (input ContainerRegistryUsernameSpec) Description() string {
	return `
  The container registry username.
`
}

func (input ContainerRegistryUsernameSpec) Link() string {
	return "container"
}

const ContainerRegistryPassword InputField = "CONTAINER_REGISTRY_PASSWORD"

type ContainerRegistryPasswordSpec struct{}

func (input ContainerRegistryPasswordSpec) Description() string {
	return `
  The container registry password.	
`
}

func (input ContainerRegistryPasswordSpec) Link() string {
	return "container"
}

const KubernetesStagingKubeconfig InputField = "KUBERNETES_STAGING_KUBECONFIG"

type KubernetesStagingKubeconfigSpec struct{}

func (input KubernetesStagingKubeconfigSpec) Description() string {
	return `
  The kubeconfig for a kubernetes cluster used for staging applications before release. Use X509 certificates with an 
  appropriate service account and rbac roles with accees to the kubernetes cluster to avoid the need for proprietary 
  authentication drivers for.
`
}

func (input KubernetesStagingKubeconfigSpec) Link() string {
	return "container"
}

const KubernetesNamespace InputField = "KUBERNETES_NAMESPACE"

type KubernetesNamespaceSpec struct{}

func (input KubernetesNamespaceSpec) Description() string {
	return `
  The namespace in the kubernetes cluster where the application will be deployed.	
`
}

func (input KubernetesNamespaceSpec) Link() string {
	return "container"
}

const SonarqubeToken InputField = "SONARQUBE_TOKEN"

type SonarqubeTokenSpec struct{}

func (input SonarqubeTokenSpec) Description() string {
	return `
  The sonarqube access token.	
`
}

func (input SonarqubeTokenSpec) Link() string {
	return "container"
}

var vars = map[string]input{
	"CONTAINER_REGISTRY":            ContainerRegistrySpec{},
	"CONTAINER_REGISTRY_USERNAME":   ContainerRegistryUsernameSpec{},
	"CONTAINER_REGISTRY_PASSWORD":   ContainerRegistryPasswordSpec{},
	"KUBERNETES_STAGING_KUBECONFIG": KubernetesStagingKubeconfigSpec{},
	"KUBERNETES_NAMESPACE":          KubernetesNamespaceSpec{},
	"SONARQUBE_TOKEN":               SonarqubeTokenSpec{},
}

func GetInputSpec(name string) input {
	return vars[name]
}
