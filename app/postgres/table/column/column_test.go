package column

import (
	"testing"

	"github.com/naysayer/schemer/api"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	Names := []struct {
		NameString string
		Expected   string
	}{
		{
			"id integer NOT NULL,",
			"id",
		},
		{
			"keyword_group_id integer,",
			"keyword_group_id",
		},
		{
			"keyword_id integer,",
			"keyword_id",
		},
		{
			`"primary" boolean DEFAULT false NOT NULL,`,
			"primary",
		},
		{
			`"position" integer,`,
			"position",
		},
		{
			"deleted boolean DEFAULT false NOT NULL,",
			"deleted",
		},
		{
			"notes text,",
			"notes",
		},
		{
			"created_at timestamp without time zone,",
			"created_at",
		},
		{
			"updated_at timestamp without time zone",
			"updated_at",
		},
	}

	for _, n := range Names {
		s, err := nameDetection(n.NameString)
		assert.NoError(t, err)
		assert.Equal(t, n.Expected, s)
	}
}

func TestType(t *testing.T) {
	Types := []struct {
		NameString string
		Expected   string
	}{
		{
			"id integer NOT NULL,",
			api.TypeInt,
		},
		{
			"keyword_group_id integer,",
			api.TypeInt,
		},
		{
			"keyword_id integer,",
			api.TypeInt,
		},
		{
			`"primary" boolean DEFAULT false NOT NULL,`,
			"bool",
		},
		{
			`"position" integer,`,
			api.TypeInt,
		},
		{
			"deleted boolean DEFAULT false NOT NULL,",
			api.TypeBool,
		},
		{
			"notes text,",
			api.TypeString,
		},
		{
			"created_at timestamp without time zone,",
			api.TypeTime,
		},
		{
			"updated_at timestamp without time zone",
			api.TypeTime,
		},
	}

	for _, ty := range Types {
		s, err := typeDetection(ty.NameString)
		assert.NoError(t, err)
		assert.Equal(t, ty.Expected, s)
	}
}
