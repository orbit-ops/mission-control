package controller

import (
	"testing"

	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/ogent"
	"github.com/orbit-ops/launchpad-core/utils"
)

func (ctrl *testAccessController) getValidMissionWithXRockets(x int) *ogent.CreateMissionReq {
	rockets := make([]uuid.UUID, 0)
	for i := 0; i < x; i++ {
		rockets = append(rockets, ctrl.createValidRocket().ID)
	}

	return &ogent.CreateMissionReq{
		Name:              "test-" + utils.RandomSeq(6),
		Rockets:           rockets,
		MinApprovers:      1,
		Duration:          1,
		PossibleApprovers: []string{"test"},
	}
}

func (ctrl *testAccessController) getValidMission() *ogent.CreateMissionReq {
	return ctrl.getValidMissionWithXRockets(1)
}

func (ctrl *testAccessController) createValidMission() *ogent.MissionCreate {
	return ctrl.createValidMissionWithXRockets(1)
}

func (ctrl *testAccessController) createValidMissionWithXRockets(x int) *ogent.MissionCreate {
	mission := ctrl.getValidMissionWithXRockets(x)
	if mc, err := ctrl.Access.CreateMission(ctrl.Context, mission); err != nil {
		ctrl.t.Fatalf("creating base valid mission: %v", err)
		return nil
	} else {
		return mc.(*ogent.MissionCreate)
	}
}

func TestMissionCreateSingleRocket(t *testing.T) {
	ctrl := newTestController(t)
	ctrl.createValidMission()
}

func TestMissionCreateMultipleRocket(t *testing.T) {
	ctrl := newTestController(t)
	ctrl.createValidMissionWithXRockets(3)
}

func TestMissionEmptyRocketsThrowError(t *testing.T) {
	ctrl := newTestController(t)

	_, err := ctrl.Access.CreateMission(ctrl.Context, &ogent.CreateMissionReq{
		Name:         "empty_rocket",
		Rockets:      []uuid.UUID{},
		MinApprovers: 1,
	})

	if err == nil {
		t.Fatal("empty rockets did not throw error")
	}
}

func TestMissionNonExistingRocketsThrowError(t *testing.T) {
	ctrl := newTestController(t)

	_, err := ctrl.Access.CreateMission(ctrl.Context, &ogent.CreateMissionReq{
		Name:              "empty_rocket",
		Rockets:           []uuid.UUID{uuid.New()},
		MinApprovers:      1,
		PossibleApprovers: []string{"test"},
	})

	if err == nil {
		t.Fatal("non-existing rockets did not throw error")
	}
}

func TestMissionEmptyApproversWhenRequiredThrowError(t *testing.T) {
	ctrl := newTestController(t)
	_, err := ctrl.Access.CreateMission(ctrl.Context, &ogent.CreateMissionReq{
		Name:              "empty_rocket",
		Rockets:           []uuid.UUID{},
		MinApprovers:      1,
		PossibleApprovers: []string{},
	})

	if err == nil {
		t.Fatal("empty possible approvers did not throw error")
	}
}

func TestMissionLessPossibleApproversThanRequiredThrowError(t *testing.T) {
	ctrl := newTestController(t)
	mission1 := ctrl.getValidMission()
	mission1.PossibleApprovers = []string{}
	if _, err := ctrl.Access.CreateMission(ctrl.Context, mission1); err == nil {
		t.Fatal("less possible approvers with 1 did not throw error")
	}

	mission2 := ctrl.getValidMission()
	mission2.MinApprovers = 2
	if _, err := ctrl.Access.CreateMission(ctrl.Context, mission2); err == nil {
		t.Fatal("less possible approvers with 2 did not throw error")
	}
}

func TestMissionDuplicateApproversThrowError(t *testing.T) {
	ctrl := newTestController(t)
	mission := ctrl.getValidMission()
	mission.MinApprovers = 1
	mission.PossibleApprovers = []string{"test", "test"}

	if _, err := ctrl.Access.CreateMission(ctrl.Context, mission); err == nil {
		t.Fatal("duplicate approvers in mission")
	}
}
