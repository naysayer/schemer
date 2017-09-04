package attr

// Attr is an interface that is used to represent a column from with a
// database and have them be able to be processed as Golang interface attributes.
// The overall goal of this package and its parents is to parse schema files and
// return golang structs. Attributes within golang structs are made up of 3
// portions which this interface must adhear to.
type Attr interface {
	Tags() string
	Title() string
	Type() string
}
