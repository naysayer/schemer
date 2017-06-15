package attr

type Attr interface {
	Name() string
	Type() string
	Tags() string
}
