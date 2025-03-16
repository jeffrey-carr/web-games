package entities

// EnvironmentDeployment respresents the deployment environment
type EnvironmentDeployment string

const (
	// EnvironmentDeploymentVariable is the environment variable representing the environment
	EnvironmentDeploymentVariable = "ENVIRONMENT"
	// DeploymentDev is the environment variable for dev deployment
	DeploymentDev EnvironmentDeployment = "dev"
	// DeploymentProd is the environment variable for prod deployment
	DeploymentProd EnvironmentDeployment = "prod"
)

// Config is the structure of the config yaml file
type Config struct {
	Environment    EnvironmentDeployment `yaml:"environment"`
	FrontendDomain string                `yaml:"frontendDomain"`
	Port           int                   `yaml:"port"`
}
