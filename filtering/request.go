package filtering

// ParseFilter parses and type-checks the filter in the provided Request.
func ParseFilter(request string, declarations *Declarations) (Filter, error) {
	if request == "" {
		return Filter{}, nil
	}
	var parser Parser
	parser.Init(request)
	parsedExpr, err := parser.Parse()
	if err != nil {
		return Filter{}, err
	}
	var checker Checker
	checker.Init(parsedExpr.Expr, parsedExpr.SourceInfo, declarations)
	checkedExpr, err := checker.Check()
	if err != nil {
		return Filter{}, err
	}
	return Filter{
		CheckedExpr: checkedExpr,
	}, nil
}
