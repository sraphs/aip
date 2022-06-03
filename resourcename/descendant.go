package resourcename

// Descendant extracts an descendant from the provided parent, using a pattern for the descendant.
func Descendant(name, pattern string) (string, bool) {
	if name == "" || pattern == "" {
		return "", false
	}
	var scName, scPattern Scanner
	scName.Init(name)
	scPattern.Init(pattern)
	for scName.Scan() {
		if !scPattern.Scan() {
			return "", false
		}
		segment := scName.Segment()
		if segment.IsWildcard() {
			return "", false // wildcards not supported in patterns
		}

		scPatternSegment := scPattern.Segment()
		if !scPatternSegment.IsVariable() && segment != scPatternSegment {
			return "", false // not a match
		}
	}

	if len(pattern) <= scPattern.end {
		return "", false
	}

	return pattern[scPattern.end+1:], true
}
