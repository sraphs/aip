package resourcename

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDescendant(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name       string
		input      string
		pattern    string
		expected   string
		expectedOK bool
	}{
		{
			name:     "empty all",
			input:    "",
			pattern:  "",
			expected: "",
		},
		{
			name:     "empty pattern",
			input:    "foo/1/bar/2",
			pattern:  "",
			expected: "",
		},

		{
			name:     "empty name",
			input:    "",
			pattern:  "foo/{foo}",
			expected: "",
		},
		{
			name:     "non-matching pattern",
			input:    "foo/1/bar/2",
			pattern:  "baz/{baz}",
			expected: "",
		},
		{
			name:     "name longer than pattern",
			input:    "foo/1/bar/2",
			pattern:  "foo/{foo}",
			expected: "",
		},
		{
			name:       "ok",
			input:      "foo/1",
			pattern:    "foo/{foo}/bar/{bar}",
			expected:   "bar/{bar}",
			expectedOK: true,
		},
		{
			name:       "ok full",
			input:      "//foo.example.com/foo/1",
			pattern:    "foo/{foo}/bar/{bar}",
			expected:   "bar/{bar}",
			expectedOK: true,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, ok := Descendant(tt.input, tt.pattern)
			assert.Equal(t, tt.expectedOK, ok)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
