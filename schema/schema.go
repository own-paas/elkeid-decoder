package schema

type Schema interface {
	Name() string
	Parse(payload []string) (err error)
}
