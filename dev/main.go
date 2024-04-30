package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
	"github.com/orbit-ops/launchpad-core/config"
	"github.com/orbit-ops/launchpad-core/ent"
	_ "github.com/orbit-ops/launchpad-core/ent/runtime"
	"github.com/orbit-ops/launchpad-core/httpserver"
	"github.com/orbit-ops/launchpad-core/providers"
	"github.com/orbit-ops/launchpad-core/providers/local"
)

func main() {
	cfg, err := config.LoadConfig(context.TODO())
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

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("running migrations: %v", err)
	}

	if err := httpserver.Start(cfg, prov, client); err != nil {
		log.Fatalf("running server: %v", err)
	}
}
