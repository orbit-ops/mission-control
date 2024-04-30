package controller

import (
	"context"
	"log"
	"os"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/orbit-ops/launchpad-core/config"
	"github.com/orbit-ops/launchpad-core/ent"
	_ "github.com/orbit-ops/launchpad-core/ent/runtime"
	"github.com/orbit-ops/launchpad-core/providers"
	"github.com/orbit-ops/launchpad-core/providers/local"
	"github.com/orbit-ops/launchpad-core/utils"
)

type testAccessController struct {
	Access  *AccessController
	Context context.Context
	t       *testing.T
}

func newTestController(t *testing.T) *testAccessController {
	os.Setenv("LAUNCHPAD_CONFIG", "../dev/launchpad.yaml")
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

	ctx := context.WithValue(context.Background(), utils.ContextUserKey{}, "test")

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("running migrations: %v", err)
	}

	ac, err := NewAccessController(prov, client)
	if err != nil {
		log.Fatalf("creating access controller: %v", err)
	}

	return &testAccessController{
		Access:  ac,
		Context: ctx,
		t:       t,
	}
}

func (ctrl *testAccessController) assert(condition bool, msg string) {
	if !condition {
		ctrl.t.Fatal(msg)
	}
}

func (ctrl *testAccessController) ok(err error, msg ...string) {
	if err != nil {
		ctrl.t.Fatalf("%s: %v", strings.Join(msg, ": "), err)
	}
}

func (ctrl *testAccessController) equals(exp, act interface{}) {
	if exp != act {
		ctrl.t.Fatal("not equal")
	}
}
