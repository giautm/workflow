package simple_test

import (
	"context"
	"testing"

	wf "github.com/giautm/workflow"
	"github.com/giautm/workflow/examples/simple"
)

func TestWorkflow(t *testing.T) {
	cfg := simple.ConfigWorkflow()

	workflow := wf.Workflow{State: simple.Init}
	workflow.SetConfig(cfg)

	if workflow.State != simple.Init {
		t.Errorf("Workflow state much be %s, received %s", simple.Init, workflow.State)
	}

	ctx := wf.NewContext(context.Background(), "approve")
	ctx.SetParam("msg", "request is approved by the first manager")
	workflow.Run(ctx)
	if workflow.State != simple.FirstApprove {
		t.Errorf("Workflow state much be %s, received %s", simple.FirstApprove, workflow.State)
	}

	ctx = wf.NewContext(context.Background(), "reject")
	ctx.SetParam("msg", "request is rejected by the second manager")
	workflow.Run(ctx)
	if workflow.State != simple.Reject {
		t.Errorf("Workflow state much be %s, received %s", simple.Reject, workflow.State)
	}

	ctx = wf.NewContext(context.Background(), "re-submit")
	ctx.SetParam("msg", "request is approved by the second manager")
	workflow.Run(ctx)
	if workflow.State != simple.SecondApprove {
		t.Errorf("Workflow state much be %s, received %s", simple.SecondApprove, workflow.State)
	}
}
