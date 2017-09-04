package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringify(t *testing.T) {
	Strings := []struct {
		Base        string
		Expected    []string
		Description string
	}{
		{
			"asdfasdf",
			[]string{"asdfasdf"},
			"basic example",
		},
		{
			`CREATE TABLE assignments (
                        updated_at timestamp without time zone);`,
			[]string{"CREATE TABLE assignments (", "updated_at timestamp without time zone);"},
			"real world example",
		},
		{
			`          CREATE TABLE assignments (
                        updated_at timestamp without time zone);   `,
			[]string{"CREATE TABLE assignments (", "updated_at timestamp without time zone);"},
			"it trims the leading and trailing spaces",
		},
	}

	for _, s := range Strings {
		assert.Equal(t, s.Expected, Seperate(s.Base), s.Description)
	}
}
