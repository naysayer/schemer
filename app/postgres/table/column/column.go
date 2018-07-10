// Package column defines a struct called column that is used for the defining
// and parsing of columns from postgres schemas.
package column

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/naysayer/schemer/api"
)

var (
	// SQL statements patterns
	pgSQLPatternInteger      = regexp.MustCompile(`^int\b`)
	pgSQLPatternBigInt       = regexp.MustCompile(`^bigint\b`)
	pgSQLPatternFloat        = regexp.MustCompile(`^double precision\b`)
	pgSQLPatternDecimalFloat = regexp.MustCompile(`^decimal\b`)
	pgSQLPatternString       = regexp.MustCompile(`^varchar\b`)
	pgSQLPatternCharString   = regexp.MustCompile(`^char\b`)
	pgSQLPatternText         = regexp.MustCompile(`^text\b`)
	pgSQLPatternTimestamp    = regexp.MustCompile(`^timestamp\b`)
	pgSQLPatternJSON         = regexp.MustCompile(`^json\b`)
	pgSQLPatternJSONB        = regexp.MustCompile(`^jsonb\b`)
	pgSQLPatternBool         = regexp.MustCompile(`^boolean\b`)
	pgSQLPatternHstore       = regexp.MustCompile(`^hstore\b`)
	pgSQLPatternDate         = regexp.MustCompile(`^date\b`)

	// Schema statments patterns
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

	// ErrBeginningOfCreate signifies that the inputted string was a create statement and not a column
	ErrBeginningOfCreate = errors.New("Beginning of create statement")
	// ErrEndingOfCreate signifies that the inputted string is the ending of a create statement from a schema file
	ErrEndingOfCreate = errors.New("End of create statement")
	// ErrUnknownColumnType indicates that the inputted string contains a column of an unknown data type
	ErrUnknownColumnType = errors.New("Unknown column type")
)

// Column conforms to the attr.Attr interface, and is used to represent
// the colums that are within a table's create statement from a postgres schema.
type Column struct {
	Name           string
	Classification string
}

// Tags returns a string that represents a db tag for a golang struct attribute
func (c *Column) Tags() string {
	return fmt.Sprintf(api.DbStructTag, c.Name)
}

// Title returns a string that represents the name of an attribute for a golang struct
func (c *Column) Title() string {
	return strings.Title(c.Name)
}

// Type returns a string that represents data type of an attribute for a golang struct
func (c *Column) Type() string {
	return c.Classification
}

// New returns a new pointer to a Column from an inputted string. The string
// argument is ideally a line from within a create statment of a postgres schema
func New(s string, sql bool) (*Column, error) {
	name, err := nameDetection(s)
	if err != nil {
		return nil, err
	}

	t, err := typeDetection(s, sql)
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

func typeDetection(s string, sql bool) (string, error) {
	fn := func(string) (string, error) {
		sep := strings.Split(s, " ")
		withoutName := append(sep[:0], sep[1:]...)
		remaining := strings.Join(withoutName, " ")

		if sql {
			return columnTypeSQL(remaining)
		}
		return columnType(remaining)
	}
	return detection(s, fn)
}

func columnType(s string) (string, error) {
	s = strings.ToLower(s)
	switch {
	case pgPatternInteger.MatchString(s):
		return api.TypeInt, nil
	case pgPatternFloat.MatchString(s):
		return api.TypeInt, nil
	case pgPatternString.MatchString(s):
		return api.TypeString, nil
	case pgPatternText.MatchString(s):
		return api.TypeString, nil
	case pgPatternTimestamp.MatchString(s):
		return api.TypeTime, nil
	case pgPatternJSON.MatchString(s):
		return api.TypeJSONText, nil
	case pgPatternJSONB.MatchString(s):
		return api.TypeJSONText, nil
	case pgPatternBool.MatchString(s):
		return api.TypeBool, nil
	case pgPatternHstore.MatchString(s):
		return api.TypeJSONText, nil
	case pgPatternDate.MatchString(s):
		return api.TypeTime, nil
	}

	return "", ErrUnknownColumnType
}

func columnTypeSQL(s string) (string, error) {
	s = strings.ToLower(s)
	switch {
	case pgSQLPatternInteger.MatchString(s) || pgSQLPatternBigInt.MatchString(s):
		return api.TypeInt, nil
	case pgSQLPatternFloat.MatchString(s):
		return api.TypeFloat, nil
	case pgSQLPatternDecimalFloat.MatchString(s):
		return api.TypeFloat, nil
	case pgSQLPatternString.MatchString(s) || pgSQLPatternCharString.MatchString(s) || pgSQLPatternText.MatchString(s):
		return api.TypeString, nil
	case pgSQLPatternTimestamp.MatchString(s):
		return api.TypeTime, nil
	case pgSQLPatternJSON.MatchString(s):
		return api.TypeJSONText, nil
	case pgSQLPatternJSONB.MatchString(s):
		return api.TypeJSONText, nil
	case pgSQLPatternBool.MatchString(s):
		return api.TypeBool, nil
	case pgSQLPatternHstore.MatchString(s):
		return api.TypeJSONText, nil
	case pgSQLPatternDate.MatchString(s):
		return api.TypeTime, nil
	}

	return "UNKNOWN", nil
	// return "", ErrUnknownColumnType
}
