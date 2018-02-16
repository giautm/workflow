package workflow

type Satisfier interface {
	IsSatisfy(ctx Context) bool
}

// SatisfierBool is fast way to use a bool constant as Satisfier
type SatisfierBool bool

// IsSatisfy implements the IsSatisfy method of the Satisfier.
func (s SatisfierBool) IsSatisfy(Context) bool {
	return bool(s)
}

// SatisfierFunc is fast way to use a func as Satisfier
type SatisfierFunc func(ctx Context) bool

// IsSatisfy implements the IsSatisfy method of the Satisfier.
func (f SatisfierFunc) IsSatisfy(ctx Context) bool {
	return f(ctx)
}
