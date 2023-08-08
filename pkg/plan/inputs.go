package plan

type InputField string

const (
	ContainerRegistry           InputField = "ContainerRegistry"
	ContainerRegistryUsername   InputField = "ContainerRegistryUsername"
	ContainerRegistryPassword   InputField = "ContainerRegistryPassword"
	KubernetesStagingKubeconfig InputField = "KubernetesStagingKubeconfig"
	KubernetesNamespace         InputField = "KubernetesNamespace"
	SonarqubeToken              InputField = "SonarqubeToken"
)
