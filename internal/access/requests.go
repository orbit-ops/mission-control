package access

import (
	"context"
	"fmt"
	"time"

	"github.com/orbit-ops/mission-control/ent"
	"github.com/orbit-ops/mission-control/ent/access"
	"github.com/orbit-ops/mission-control/ent/approval"
	"github.com/orbit-ops/mission-control/ent/ogent"
	"github.com/orbit-ops/mission-control/ent/request"
)

type Controller struct {
	*ogent.OgentHandler
	client *ent.Client
}

func (c *Controller) CreateApproval(ctx context.Context, req *ent.Request, approver string) error {
	approved, err := c.client.Approval.Query().Where(approval.RequestIDEQ(req.ID)).Count(ctx)

	r, err := c.client.Request.Query().Where(request.IDEQ(req.ID)).WithMission().First(ctx)

	_, err = c.client.Approval.Create().
		SetApproved(true).
		SetApprovedTime(time.Now()).
		SetPerson(approver).
		Save(ctx)

	if err != nil {
		return nil
	}

	if approved+1 < r.Edges.Mission.MinApprovers {
		return nil
	}

	return c.CreateAccess(ctx, req)
	// audit
	// are all approvals done?
	// 		CreateAccess()
}

func (c *Controller) RemoveApproval(ctx context.Context, req *ent.Request, approver string) error {
	exists, err := c.client.Access.Query().Where(access.RequestIDEQ(req.ID)).Exist(ctx)
	if err != nil {
		return fmt.Errorf("checking access already provisioned: %w", err)
	}
	if exists {
		return fmt.Errorf("cannot revoke approval, access provisioned")
	}

	_, err = c.client.Approval.Delete().Where(approval.PersonEQ(approver), approval.ApprovedEQ(true), approval.RequestIDEQ(req.ID)).Exec(ctx)
	if err != nil {
		return fmt.Errorf("executing approval removal: %w", err)
	}

	return nil
}

func (c *Controller) CancelAccess(ctx context.Context, acc *ent.Access) error {
	err := acc.Update().SetRolledBack(true).SetRollbackTime(time.Now()).Exec(ctx)
	if err != nil {
		return err
	}

	return c.RemoveAccess(ctx, acc)
}

func (c *Controller) CreateAccess(ctx context.Context, req *ent.Request) error {
	// audit
	// scheduleDeletion()

	return nil
}

func (c *Controller) RemoveAccess(ctx context.Context, acc *ent.Access) error {
	return nil
	// audit
}

func (c *Controller) scheduleDeletion(ctx context.Context, appr *ent.Approval) {

}
