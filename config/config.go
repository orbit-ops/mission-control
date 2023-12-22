package config

import (
	"fmt"

	"github.com/orbit-ops/launchpad-core/utils"
	"github.com/spf13/viper"
	resolvers "github.com/tiagoposse/go-secret-resolvers"
)

type ProviderType string

const (
	ProviderKubernetes ProviderType = "kubernetes"
	ProviderAws        ProviderType = "aws"
	ProviderDocker     ProviderType = "docker"
	ProviderLocal      ProviderType = "local"
)

type ProviderConfig struct {
	Type       ProviderType              `yaml:"type"`
	Executable string                    `yaml:"executable"`
	Kubernetes *ProviderKubernetesConfig `yaml:"kubernetes"`
}

type ProviderKubernetesConfig struct {
	JobNamespace string `yaml:"jobNamespace"`
}

type SessionsConfig struct {
	SessionKey    *resolvers.ResolverField `yaml:"key"`
	JwtExpiration utils.Duration           `yaml:"expiration"`
}

type Config struct {
	DevMode  bool            `yaml:"dev"`
	Version  string          `yaml:"version"`
	Provider ProviderConfig  `yaml:"provider"`
	ApiUrl   string          `yaml:"apiUrl"`
	Sessions *SessionsConfig `yaml:"sessions"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("launchpad")        // name of config file (without extension)
	viper.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/launchpad/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.launchpad") // call multiple times to add many search paths
	viper.AddConfigPath(".")                // optionally look for config in the working directory
	viper.SetEnvPrefix("launchpad")
	viper.AutomaticEnv()
	viper.SetDefault("provider.type", ProviderKubernetes)
	viper.SetDefault("apiUrl", "http://localhost:9000")
	viper.SetDefault("dev", false)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var cfg *Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return cfg, nil
}
