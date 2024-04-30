package controller

import (
	"context"
	"fmt"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/hook"
	"github.com/orbit-ops/launchpad-core/ent/mission"
)

func (ac *AccessController) setAccessHooks() {
	// do not allow updates
	ac.client.Request.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.RequestFunc(func(ctx context.Context, am *ent.RequestMutation) (ent.Value, error) {
			return nil, fmt.Errorf("requests cannot be updated")
		})
	}, ent.OpUpdate))

	// require valid, existing missions
	ac.client.Request.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.RequestFunc(func(ctx context.Context, rm *ent.RequestMutation) (ent.Value, error) {
			missionID, _ := rm.MissionID()
			if exists, err := rm.Client().Mission.Query().Where(mission.IDEQ(missionID)).Exist(ctx); err != nil {
				return nil, err
			} else if !exists {
				return nil, fmt.Errorf("mission %s does not exist", missionID.String())
			}
			return next.Mutate(ctx, rm)
		})
	}, ent.OpCreate))
}
