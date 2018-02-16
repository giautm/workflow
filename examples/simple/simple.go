package simple

import (
	wf "github.com/giautm/workflow"
)

var workflow wf.Workflow

const (
	Init          = wf.State("init")
	FirstApprove  = wf.State("first_approve")
	SecondApprove = wf.State("second_approve")
	Approve       = wf.State("approve")
	Reject        = wf.State("reject")
)

type A struct{}

func (a *A) Approve(wf.Context) error {
	return nil
}

var a A

func ConfigWorkflow() wf.Config {
	config := wf.NewConfig(Init, wf.ConfigDesc{
		Init: {
			Activities: wf.Activity{create, wf.ActivityFunc(a.Approve)},
			Transitions: wf.Transitions{
				{
					ToState: FirstApprove,
					Trigger: "approve",
				},
			},
		},
		FirstApprove: {
			Activities: wf.Activity{mail},
			Transitions: wf.Transitions{
				{
					ToState: SecondApprove,
					Trigger: "approve",
				},
				{
					ToState: Reject,
					Trigger: "reject",
				},
			},
		},
		SecondApprove: {
			Activities: wf.Activity{mail, complete},
		},
		Reject: {
			Activities: wf.Activity{mail},
			Transitions: wf.Transitions{
				{
					ToState: SecondApprove,
					Trigger: "re-submit",
				},
			},
		},
	})

	return config
}
