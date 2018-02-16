package workflow

type Executable interface {
	Execute(ctx Context) error
}

type Activity []Executable

func (e Activity) Execute(ctx Context) error {
	for _, a := range e {
		if err := a.Execute(ctx); err != nil {
			return err
		}
	}

	return nil
}

// ActivityFunc is a function that can be used as a Executable.
type ActivityFunc func(Context) error

// Execute implements the Execute method of the Executable.
func (a ActivityFunc) Execute(ctx Context) error {
	return a(ctx)
}
