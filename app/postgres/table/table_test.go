package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	cases := []struct {
		NameString  string
		Expected    string
		Description string
	}{
		{
			"CREATE TABLE admins (",
			"admins",
			"basic example",
		},
		{
			"CREATE TABLE strFirstNameTableName (",
			"strFirstNameTableName",
			"hungarian notation",
		},
		{
			"SOME RANDOM STRING",
			"SOME RANDOM STRING",
			"pattern not present",
		},
	}

	for _, c := range cases {
		s := Name(c.NameString)
		assert.Equal(t, c.Expected, s)
	}
}
