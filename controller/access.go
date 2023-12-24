package controller

import (
	"context"
	"fmt"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/hook"
)

func (ac *AccessController) setRequestHooks() {
	ac.client.Access.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AccessFunc(func(ctx context.Context, am *ent.AccessMutation) (ent.Value, error) {
			return nil, fmt.Errorf("accesses cannot be deleted")
		})
	}, ent.OpDelete))
}
