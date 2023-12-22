package main

import (
	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
	"github.com/orbit-ops/launchpad-core/config"
	"github.com/orbit-ops/launchpad-core/controller"
	"github.com/orbit-ops/launchpad-core/providers"
	"github.com/orbit-ops/launchpad-core/providers/local"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	var prov providers.Provider
	provCfg := &providers.ProviderConfig{
		ApiUrl: cfg.ApiUrl,
	}

	prov, err = local.NewLocalProvider(provCfg)
	if err != nil {
		log.Fatalf("Initializing %s provider: %v\n", cfg.Provider.Type, err)
	}

	_, err = controller.NewAccessController(prov)
	if err != nil {
		log.Fatal(err)
	}
}
