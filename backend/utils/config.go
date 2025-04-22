package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

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
	FullCertPath   string                `yaml:"fullCertPath"`
	PrivateKeyPath string                `yaml:"privateKeyPath"`

	WordLadder WordLadderConfig `yaml:"wordLadder"`
}

// WordLadderConfig is the configuration for the WordLadder game
type WordLadderConfig struct {
	// MaxServers is the maxmium number of servers (concurrent games)
	MaxServers int `yaml:"maxServers"`
	// MaxPlayersPerServer is the maximum number of players per game
	MaxPlayersPerServer int `yaml:"maxPlayersPerServer"`
}

// ReadConfig reads the provided configuration file
func ReadConfig(filename string) (Config, error) {
	// Load file
	file, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	// Read file into entity
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
