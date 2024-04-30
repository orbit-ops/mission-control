package controller

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/approval"
	"github.com/orbit-ops/launchpad-core/ent/mission"
	"github.com/orbit-ops/launchpad-core/ent/request"
)

func (ctrl *testAccessController) getValidAccess() *ent.AccessCreate {
	appr := ctrl.createValidApproval()

	req, _ := ctrl.Access.client.Request.Query().WithMission().Where(request.HasApprovalsWith(approval.IDEQ(appr.ID))).First(ctrl.Context)

	from := time.Now()
	to := from.Add(time.Duration(req.Edges.Mission.Duration) * time.Minute)
	return ctrl.Access.client.Access.Create().
		SetExpiration(to).
		SetRequestID(req.ID).
		SetStartTime(from)
}

func (ctrl *testAccessController) createValidAccess() *ent.Access {
	resp, err := ctrl.getValidAccess().Save(ctrl.Context)
	if err != nil {
		ctrl.t.Fatal(err)
	}
	return resp
}

func TestAccessCreate(t *testing.T) {
	ctrl := newTestController(t)
	ctrl.createValidAccess()
}

func TestAccessCreateWithoutEnoughApprovalsThrowError(t *testing.T) {
	ctrl := newTestController(t)
	req := ctrl.createValidRequest()

	m, _ := ctrl.Access.client.Mission.Query().Where(mission.HasRequestsWith(request.IDEQ(req.ID))).First(ctrl.Context)
	from := time.Now()
	to := from.Add(time.Duration(m.Duration) * time.Minute)

	if _, err := ctrl.Access.client.Access.Create().
		SetExpiration(to).
		SetRequestID(req.ID).
		SetStartTime(from).Save(ctrl.Context); err == nil {
		t.Fatal("access is being created without enough approvals being given")
	}
}

func TestAccessWithNonExistingRequestThrowError(t *testing.T) {
	ctrl := newTestController(t)
	acc := ctrl.getValidAccess()

	if _, err := acc.SetRequestID(uuid.UUID{}).Save(ctrl.Context); err == nil {
		t.Fatal("access with unknown request no error")
	}
}

// func TestAccessUpdateAccessThrowError(t *testing.T) {
// 	ctrl := newTestController(t)
// 	acc := ctrl.createValidAccess()

// }
