package filtering

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/sraphs/aip/filtering/internal/testdata"
)

func TestChecker(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		filter        string
		declarations  []DeclarationOption
		errorContains string
	}{
		{
			filter: `single_string=1`,
			declarations: []DeclarationOption{
				DeclareStandardFunctions(),
				DeclareProtoMessage(new(testdata.TestAllTypes)),
			},
		},
	} {
		tt := tt
		t.Run(tt.filter, func(t *testing.T) {
			t.Parallel()
			var parser Parser
			parser.Init(tt.filter)
			parsedExpr, err := parser.Parse()
			if err != nil && tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
				return
			}
			assert.NilError(t, err)
			declarations, err := NewDeclarations(tt.declarations...)
			if err != nil && tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
				return
			}
			assert.NilError(t, err)
			var checker Checker
			checker.Init(parsedExpr.Expr, parsedExpr.SourceInfo, declarations)
			checkedExpr, err := checker.Check()
			if tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
				return
			}
			assert.NilError(t, err)
			assert.Assert(t, checkedExpr != nil)
		})
	}
}
