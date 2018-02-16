package workflow

type Transition interface {
	FromState() State
	NextState() State
	IsSatisfy(ctx Context) bool
}

func NewTransition(fromState, nextState State) Transition {
	return NewTransitionWithSatisfier(fromState, nextState, SatisfierBool(true))
}

func NewTransitionWithSatisfier(
	fromState, nextState State,
	satisfier Satisfier,
) Transition {
	return &transition{
		from:      fromState,
		next:      nextState,
		satisfier: satisfier,
	}
}

type transition struct {
	from, next State
	satisfier  Satisfier
}

// FromState returns current state of transition
func (t *transition) FromState() State {
	return t.from
}

// NextState returns next state of transition
func (t *transition) NextState() State {
	return t.next
}

// IsSatisfy implements the IsSatisfy method of the Satisfier.
func (t *transition) IsSatisfy(ctx Context) bool {
	return t.satisfier.IsSatisfy(ctx)
}
