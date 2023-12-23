package config

import (
	"context"
	"fmt"
	"os"

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
	Type       ProviderType              `mapstructure:"type"`
	Executable string                    `mapstructure:"executable"`
	Kubernetes *ProviderKubernetesConfig `mapstructure:"kubernetes"`
}

type ProviderKubernetesConfig struct {
	JobNamespace string `mapstructure:"jobNamespace"`
}

type SessionsConfig struct {
	SessionKey    *resolvers.ResolverField `mapstructure:"key"`
	JwtExpiration utils.Duration           `mapstructure:"expiration"`
}

type Config struct {
	DevMode  bool            `mapstructure:"dev"`
	Version  string          `mapstructure:"version"`
	Provider ProviderConfig  `mapstructure:"provider"`
	ApiUrl   string          `mapstructure:"apiUrl"`
	Sessions *SessionsConfig `mapstructure:"sessions"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	viper.SetConfigName("launchpad")        // name of config file (without extension)
	viper.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/launchpad/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.launchpad") // call multiple times to add many search paths
	viper.AddConfigPath(".")                // optionally look for config in the working directory
	// viper.AddConfigPath(os.Getenv("LAUNCHPAD_CONFIG")) // optionally look for config in the working directory
	if val, ok := os.LookupEnv("LAUNCHPAD_CONFIG"); ok {
		viper.SetConfigFile(val)
	}
	viper.SetEnvPrefix("launchpad")
	viper.AutomaticEnv()
	viper.SetDefault("provider.type", ProviderKubernetes)
	viper.SetDefault("apiUrl", "localhost:9000")
	viper.SetDefault("dev", false)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var cfg *Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	resolv := resolvers.NewResolver()
	if err := resolv.Resolve(ctx, cfg.Sessions.SessionKey); err != nil {
		return nil, fmt.Errorf("resolving session key: %w", err)
	}
	return cfg, nil
}
