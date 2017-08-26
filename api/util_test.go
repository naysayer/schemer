package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringify(t *testing.T) {
	Strings := []struct {
		Base     string
		Expected []string
	}{
		{
			"asdfasdf",
			[]string{"asdfasdf"},
		},
		{
			`CREATE TABLE assignments (
                        updated_at timestamp without time zone);`,
			[]string{"CREATE TABLE assignments (", "updated_at timestamp without time zone);"},
		},
	}

	for _, s := range Strings {
		actual := Seperate(s.Base)
		assert.Equal(t, s.Expected, actual)
	}
}
