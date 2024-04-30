package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ogen-go/ogen/middleware"
	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/access"
	"github.com/orbit-ops/launchpad-core/ent/approval"
	"github.com/orbit-ops/launchpad-core/ent/hook"
	"github.com/orbit-ops/launchpad-core/ent/ogent"
	"github.com/orbit-ops/launchpad-core/ent/request"
)

func (ac *AccessController) setRequestHooks() {
	// prevent deletion
	ac.client.Access.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AccessFunc(func(ctx context.Context, am *ent.AccessMutation) (ent.Value, error) {
			amID, _ := am.ID()
			existing, _ := am.Client().Access.Query().Where(access.IDEQ(amID)).First(ctx)
			if existing.RolledBack {
				return nil, fmt.Errorf("access has been rolled back and cannot be further modified")
			}

			rback, _ := am.RolledBack()
			if !rback {
				return nil, fmt.Errorf("accesses cannot be modified")
			}

			am.SetRollbackTime(time.Now())
			return next.Mutate(ctx, am)
		})
	}, ent.OpUpdate))

	// prevent access creation if request does not have enough approvals
	ac.client.Access.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AccessFunc(func(ctx context.Context, am *ent.AccessMutation) (ent.Value, error) {
			reqID, _ := am.RequestID()

			req, err := am.Client().Request.Query().WithMission().WithApprovals().Where(request.IDEQ(reqID)).First(ctx)
			if err != nil {
				return nil, err
			}

			if req.Edges.Mission.MinApprovers > len(req.Edges.Approvals) {
				return nil, fmt.Errorf("request does not have enough approvals")
			}
			return next.Mutate(ctx, am)
		})
	}, ent.OpCreate))

	ac.client.Access.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AccessFunc(func(ctx context.Context, am *ent.AccessMutation) (ent.Value, error) {
			return nil, fmt.Errorf("accesses cannot be deleted")
		})
	}, ent.OpDelete))
}

func (ac *AccessController) setAccessMiddleware() middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		resp, err := next(req)

		switch req.OperationID {
		case "createApproval":

		case "updateApproval":
		}

		return resp, err
	}
}

// CreateApproval handles POST /approvals requests.
func (h *AccessController) CreateApproval(ctx context.Context, req *ogent.CreateApprovalReq) (ogent.CreateApprovalRes, error) {
	b := h.client.Approval.Create()
	// Add all fields.
	b.SetPerson(req.Person)
	b.SetApprovedTime(req.ApprovedTime)
	b.SetApproved(req.Approved)
	b.SetRevoked(req.Revoked)
	if v, ok := req.RevokedTime.Get(); ok {
		b.SetRevokedTime(v)
	}
	// Add all edges.
	b.SetRequestID(req.Request)
	if v, ok := req.Access.Get(); ok {
		b.SetAccessID(v)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &ogent.R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &ogent.R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Approval.Query().Where(approval.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return ogent.NewApprovalCreate(e), nil
}
