package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/access"
	"github.com/orbit-ops/launchpad-core/ent/approval"
	"github.com/orbit-ops/launchpad-core/ent/hook"
	"github.com/orbit-ops/launchpad-core/ent/request"
)

func (ac *AccessController) setApprovalHooks() {
	// do not allow duplicate <request, person> approvals
	ac.client.Approval.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.ApprovalFunc(func(ctx context.Context, am *ent.ApprovalMutation) (ent.Value, error) {
			reqID, ok := am.RequestID()
			if !ok {
				return nil, errors.New("request ID not set")
			}
			person, ok := am.Person()
			if !ok {
				return nil, errors.New("person not set")
			}
			exists, err := am.Client().Approval.Query().Where(approval.HasRequestWith(request.IDEQ(reqID)), approval.PersonEQ(person)).Exist(ctx)
			if err != nil {
				return nil, fmt.Errorf("checking approval exists: %w", err)
			}
			if exists {
				return nil, fmt.Errorf("%s already approved %s", person, reqID)
			}

			// approval.RequestIDEQ(reqID)
			return next.Mutate(ctx, am)
		})
	}, ent.OpCreate))

	// create access if all approvals given
	ac.client.Approval.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.ApprovalFunc(func(ctx context.Context, am *ent.ApprovalMutation) (ent.Value, error) {
			reqID, ok := am.RequestID()
			if !ok {
				return nil, errors.New("request ID not set")
			}

			req, err := am.Client().Request.Query().
				Where(request.IDEQ(reqID)).
				WithMission().
				WithApprovals().First(ctx)
			if err != nil {
				return nil, fmt.Errorf("checking approval exists: %w", err)
			}
			if len(req.Edges.Approvals) >= req.Edges.Mission.MinApprovers {
				go ac.CreateAccess(ctx, req.Edges.Mission, req)
			}

			return next.Mutate(ctx, am)
		})
	}, ent.OpCreate))

	ac.client.Approval.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.ApprovalFunc(func(ctx context.Context, am *ent.ApprovalMutation) (ent.Value, error) {
			reqID, ok := am.RequestID()
			if !ok {
				return nil, errors.New("request ID not set")
			}

			exists, err := am.Client().Access.Query().
				Where(access.RequestIDEQ(reqID)).Exist(ctx)
			if err != nil {
				return nil, fmt.Errorf("checking access exists: %w", err)
			}
			if exists {
				return nil, fmt.Errorf("cannot update approval once access provisioned")
			}

			return next.Mutate(ctx, am)
		})
	}, ent.OpUpdate))
}
