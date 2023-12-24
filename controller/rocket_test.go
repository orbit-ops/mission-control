package controller

import (
	"testing"

	"github.com/orbit-ops/launchpad-core/ent/ogent"
	"github.com/orbit-ops/launchpad-core/utils"
)

func (*testAccessController) getValidRocket() *ogent.CreateRocketReq {
	return &ogent.CreateRocketReq{
		Name: "test-" + utils.RandomSeq(6),
	}
}

func (ctrl *testAccessController) createValidRocket() *ogent.RocketCreate {
	resp, err := ctrl.Access.CreateRocket(ctrl.Context, ctrl.getValidRocket())
	if err != nil {
		ctrl.t.Fatalf("creating base rocket: %v", err)
	}
	return resp.(*ogent.RocketCreate)
}

func TestRocketCreate(t *testing.T) {
	ctrl := newTestController(t)
	ctrl.createValidRocket()
}

func TestRocketNameIsEmptyError(t *testing.T) {
	ctrl := newTestController(t)
	_, err := ctrl.Access.CreateRocket(ctrl.Context, &ogent.CreateRocketReq{})

	if err == nil {
		t.Fatal("Rockets are accepting empty names")
	}
}
