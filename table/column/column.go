package column

import (
	"errors"
	"regexp"
)

var (
	CreateTableRegex = regexp.MustCompile("^CREATE TABLE")
	FirstWordRegex   = regexp.MustCompile("(?:^|(?:[.!?]\\s))(\\w+)\\s")
)

var (
	PgIntegerRegex   = regexp.MustCompile("^integer\\b")
	PgFloatRegex     = regexp.MustCompile("^double precision\\b")
	PgStringRegex    = regexp.MustCompile("^character varying\\b")
	PgTextRegex      = regexp.MustCompile("^text\\b")
	PgTimestampRegex = regexp.MustCompile("^timestamp\\b")
	PgJsonRegex      = regexp.MustCompile("^json\\b")
	PgJsonBRegex     = regexp.MustCompile("^jsonb\\b")
	PgBoolRegex      = regexp.MustCompile("^boolean\\b")
	PgHstoreRegex    = regexp.MustCompile("^hstore\\b")
	PgDateRegex      = regexp.MustCompile("^date\\b")
)

var (
	BeginningOfCreateError = errors.New("Beginning of create statement")
	EndingOfCreateError    = errors.New("End of create statement")

	UnknownColumnType = errors.New("Unknown column type")
)

type Column struct {
	Name string
	Type string
}

func (c Column) Populate(s string) (*Column, error) {
	name, err := Name(s)
	if err != nil {
		return &Column{}, err
	}

	t, err := Type(s)
	if err != nil {
		return &Column{}, err
	}

	return &Column{Name: name, Type: t}, nil
}

func guard(s string) error {
	if s == ");" {
		return EndingOfCreateError
	}
	if CreateTableRegex.MatchString(s) {
		return BeginningOfCreateError
	}
	return nil

}
func Name(s string) (string, error) {
	err := guard(s)
	if err != nil {
		return "", err
	}
	return FirstWordRegex.FindString(s), nil
}

func Type(s string) (string, error) {
	err := guard(s)
	if err != nil {
		return "", err
	}

	remaining := FirstWordRegex.ReplaceAllString(s, "") // remove first word
	return columnType(remaining)
}

func columnType(s string) (string, error) {
	switch {
	case PgIntegerRegex.MatchString(s):
		return "int", nil
	case PgFloatRegex.MatchString(s):
		return "float", nil
	case PgStringRegex.MatchString(s):
		return "string", nil
	case PgTextRegex.MatchString(s):
		return "string", nil
	case PgTimestampRegex.MatchString(s):
		return "time.Time", nil
	case PgJsonRegex.MatchString(s):
		return "types.JSONText", nil
	case PgJsonBRegex.MatchString(s):
		return "types.JSONText", nil
	case PgBoolRegex.MatchString(s):
		return "bool", nil
	case PgHstoreRegex.MatchString(s):
		return "types.JSONText", nil // not sure this is going to work
	case PgDateRegex.MatchString(s):
		return "time.Time", nil // not sure this is going to work
	}

	return "", UnknownColumnType
}
