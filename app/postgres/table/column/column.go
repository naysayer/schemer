// Package column defines a struct called column that is used for the defining
// and parsing of columns from postgres schemas.
package column

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	pgPatternCreateTable = regexp.MustCompile(`^CREATE TABLE`)
	pgPatternwrapped     = regexp.MustCompile(`\"`)
	pgPatternInteger     = regexp.MustCompile(`^integer\b`)
	pgPatternFloat       = regexp.MustCompile(`^double precision\b`)
	pgPatternString      = regexp.MustCompile(`^character varying\b`)
	pgPatternText        = regexp.MustCompile(`^text\b`)
	pgPatternTimestamp   = regexp.MustCompile(`^timestamp\b`)
	pgPatternJSON        = regexp.MustCompile(`^json\b`)
	pgPatternJSONB       = regexp.MustCompile(`^jsonb\b`)
	pgPatternBool        = regexp.MustCompile(`^boolean\b`)
	pgPatternHstore      = regexp.MustCompile(`^hstore\b`)
	pgPatternDate        = regexp.MustCompile(`^date\b`)

	endOfCreate = ");"
	dbStructTag = "`db:\"%v\"`" // very much like the golang types, this needs to be in its own package so it can be usable across the program

	// ErrBeginningOfCreate signifies that the inputted string was a create statement and not a column
	ErrBeginningOfCreate = errors.New("Beginning of create statement")
	// ErrEndingOfCreate signifies that the inputted string is the ending of a create statement from a schema file
	ErrEndingOfCreate = errors.New("End of create statement")
	// ErrUnknownColumnType indicates that the inputted string contains a column of an unknown data type
	ErrUnknownColumnType = errors.New("Unknown column type")
)

type Column struct {
	Name           string
	Classification string
}

func (c *Column) Tags() string {
	return fmt.Sprintf(dbStructTag, c.Name)
}
func (c *Column) Title() string {
	return strings.Title(c.Name)
}
func (c *Column) Type() string {
	return c.Classification
}

func New(s string) (*Column, error) {
	name, err := nameDetection(s)
	if err != nil {
		return nil, err
	}

	t, err := typeDetection(s)
	if err != nil {
		return nil, err
	}

	return &Column{Name: name, Classification: t}, nil
}

func detection(s string, fn func(string) (string, error)) (string, error) {
	if s == endOfCreate {
		return "", ErrEndingOfCreate
	}
	if pgPatternCreateTable.MatchString(s) {
		return "", ErrBeginningOfCreate
	}
	return fn(s)
}

func nameDetection(s string) (string, error) {
	fn := func(string) (string, error) {
		sep := strings.Split(s, " ")                           // seperates line into array by spaces
		rawName := fmt.Sprintf("%v", sep[0])                   // interplate value in case it is nil
		name := pgPatternwrapped.ReplaceAllString(rawName, "") // In the event the name of the column is wrapped with quotes we strip it out
		return name, nil
	}
	return detection(s, fn)
}

func typeDetection(s string) (string, error) {
	fn := func(string) (string, error) {
		sep := strings.Split(s, " ")
		withoutName := append(sep[:0], sep[1:]...)
		remaining := strings.Join(withoutName, " ")

		return columnType(remaining)
	}
	return detection(s, fn)
}

// TODO: these returned strings are golang types either from this package or
// a 3rd party package like sqlx. They are static across databasese as these
// represent their golang counterparts so they should be extracted into a
// package where they can be shared across this program.
func columnType(s string) (string, error) {
	switch {
	case pgPatternInteger.MatchString(s):
		return "int", nil
	case pgPatternFloat.MatchString(s):
		return "float", nil
	case pgPatternString.MatchString(s):
		return "string", nil
	case pgPatternText.MatchString(s):
		return "string", nil
	case pgPatternTimestamp.MatchString(s):
		return "time.Time", nil
	case pgPatternJSON.MatchString(s):
		return "types.JSONText", nil
	case pgPatternJSONB.MatchString(s):
		return "types.JSONText", nil
	case pgPatternBool.MatchString(s):
		return "bool", nil
	case pgPatternHstore.MatchString(s):
		return "types.JSONText", nil
	case pgPatternDate.MatchString(s):
		return "time.Time", nil
	}

	return "", ErrUnknownColumnType
}
