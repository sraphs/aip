package ordering

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestUnmarshal(t *testing.T) {
	for _, tt := range []struct {
		orderBy       string
		expected      OrderBy
		errorContains string
	}{
		{
			orderBy:  "",
			expected: OrderBy{},
		},

		{
			orderBy: "foo desc, bar",
			expected: OrderBy{
				Fields: []Field{
					{Path: "foo", Desc: true},
					{Path: "bar"},
				},
			},
		},

		{
			orderBy: "foo.bar",
			expected: OrderBy{
				Fields: []Field{
					{Path: "foo.bar"},
				},
			},
		},

		{
			orderBy: " foo , bar desc ",
			expected: OrderBy{
				Fields: []Field{
					{Path: "foo"},
					{Path: "bar", Desc: true},
				},
			},
		},

		{orderBy: "foo,", errorContains: "invalid format"},
		{orderBy: ",", errorContains: "invalid "},
		{orderBy: ",foo", errorContains: "invalid format"},
		{orderBy: "foo/bar", errorContains: "invalid character '/'"},
		{orderBy: "foo bar", errorContains: "invalid format"},
	} {
		tt := tt
		t.Run(tt.orderBy, func(t *testing.T) {
			t.Parallel()
			actual, err := Unmarshal(tt.orderBy)
			if tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
			} else {
				assert.NilError(t, err)
				assert.DeepEqual(t, tt.expected, actual)
			}
		})
	}
}

func TestField_SubFields(t *testing.T) {
	for _, tt := range []struct {
		name     string
		field    Field
		expected []string
	}{
		{
			name:     "empty",
			field:    Field{},
			expected: nil,
		},

		{
			name:     "single",
			field:    Field{Path: "foo"},
			expected: []string{"foo"},
		},

		{
			name:     "multiple",
			field:    Field{Path: "foo.bar"},
			expected: []string{"foo", "bar"},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.DeepEqual(t, tt.expected, tt.field.SubFields())
		})
	}
}
