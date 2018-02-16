package workflow

import (
	"errors"
	"time"

	"github.com/tikivn/github.com/satori/go.uuid"
)

// State is state of workflow
type State string

// Workflow is core model
type Workflow struct {
	Id        uuid.UUID
	State     State
	Version   int
	CreatedAt time.Time
	config    Config
}

func (w *Workflow) SetConfig(config Config) *Workflow {
	w.config = config
	return w
}

func (w *Workflow) Run(ctx Context) error {
	if w.config == nil {
		return errors.New("Workflow config is nil")
	}

	transitions := w.config.Transitions(w.State, ctx.Trigger())
	if transitions != nil {
		for _, t := range transitions {
			if t.IsSatisfy(ctx) {
				w.State = t.NextState()

				activity := w.config.Activity(w.State)
				activity.Execute(ctx)

				w.Version++
				return nil
			}
		}
	}

	return nil
}
