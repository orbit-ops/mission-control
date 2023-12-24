package controller

import (
	"context"
	"fmt"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/hook"
)

func (ac *AccessController) setAccessHooks() {
	ac.client.Request.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.RequestFunc(func(ctx context.Context, am *ent.RequestMutation) (ent.Value, error) {
			return nil, fmt.Errorf("requests cannot be updated")
		})
	}, ent.OpUpdate))
}
