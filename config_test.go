package workflow_test

import (
	"testing"

	wf "github.com/giautm/workflow"
)

func TestNewConfig(t *testing.T) {
	config := wf.NewConfig("opened", wf.ConfigDesc{
		"opened": wf.StateDesc{
			Activities: wf.Activity{
				wf.ActivityFunc(func(wf.Context) error {
					return nil
				}),
			},
			Transitions: []wf.TransitionDesc{
				wf.TransitionDesc{
					ToState: "closed",
					Trigger: "close",
				},
				wf.TransitionDesc{
					ToState: "closed",
					Trigger: "close",
					Satisfier: wf.SatisfierFunc(func(ctx wf.Context) bool {
						u, ok := ctx.Param("user")
						return ok && u == "giautm"
					}),
				},
			},
		},
	})

	if ts := config.Transitions("opened", "close"); len(ts) == 0 {
		t.Error("AAAA")
	}
}
