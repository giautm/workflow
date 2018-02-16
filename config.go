package workflow

type Config interface {
	Activity(state State) Activity
	InitState() State
	Transitions(fromState State, trigger string) []Transition
}

type ConfigDesc map[State]StateDesc

type StateDesc struct {
	Activities  Activity
	Transitions Transitions
}

type Transitions []TransitionDesc

type TransitionDesc struct {
	Satisfier Satisfier
	ToState   State
	Trigger   string
}

// NewConfig returns config for Workflow
func NewConfig(initState State, cfg ConfigDesc) Config {
	config := &config{
		initState:   initState,
		activities:  map[State]Activity{},
		transitions: map[eKey][]Transition{},
	}

	for state, desc := range cfg {
		config.activities[state] = desc.Activities

		for _, t := range desc.Transitions {
			key := eKey{t.Trigger, state}
			ts, ok := config.transitions[key]
			if !ok {
				ts = []Transition{}
			}

			if t.Satisfier == nil {
				config.transitions[key] = append(ts, NewTransition(
					state, t.ToState,
				))
			} else {
				config.transitions[key] = append(ts, NewTransitionWithSatisfier(
					state, t.ToState, t.Satisfier,
				))
			}
		}
	}

	return config
}

type config struct {
	initState   State
	activities  map[State]Activity
	transitions map[eKey][]Transition
}

var _ = Config(&config{})

func (c *config) InitState() State {
	return c.initState
}

func (c *config) Activity(state State) Activity {
	if activity, ok := c.activities[state]; ok {
		return activity
	}

	return nil
}

func (c *config) Transitions(fromState State, trigger string) []Transition {
	key := eKey{trigger, fromState}
	if transitions, ok := c.transitions[key]; ok {
		return transitions
	}

	return nil
}

// eKey is a struct key used for storing the transition map.
type eKey struct {
	trigger string
	state   State
}
