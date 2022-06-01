// Package ordering provides primitives for implementing AIP ordering.
//
// See: https://google.aip.dev/132#ordering (Standard methods: List > Ordering).
package ordering

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Unmarshal request parses the ordering field for a Request.
func Unmarshal(s string) (OrderBy, error) {
	o := OrderBy{}

	if s == "" {
		return o, nil
	}

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '_' && r != ' ' && r != ',' && r != '.' {
			return o, fmt.Errorf("unmarshal order by '%s': invalid character %s", s, strconv.QuoteRune(r))
		}
	}

	fields := strings.Split(s, ",")

	o.Fields = make([]Field, 0, len(fields))

	for _, field := range fields {
		parts := strings.Fields(field)
		switch len(parts) {
		case 1: // default ordering (ascending)
			o.Fields = append(o.Fields, Field{Path: parts[0]})
		case 2: // specific ordering
			order := parts[1]
			var desc bool
			switch order {
			case "asc":
				desc = false
			case "desc":
				desc = true
			default: // parse error
				return o, fmt.Errorf("unmarshal order by '%s': invalid format", s)
			}
			o.Fields = append(o.Fields, Field{Path: parts[0], Desc: desc})
		case 0:
			fallthrough
		default:
			return o, fmt.Errorf("unmarshal order by '%s': invalid format", s)
		}
	}
	return o, nil
}

// OrderBy represents an ordering directive.
type OrderBy struct {
	// Fields are the fields to order by.
	Fields []Field
}

// Field represents a single ordering field.
type Field struct {
	// Path is the path of the field, including subfields.
	Path string
	// Desc indicates if the ordering of the field is descending.
	Desc bool
}

// SubFields returns the individual subfields of the field path, including the top-level subfield.
//
// Subfields are specified with a . character, such as foo.bar or address.street.
func (f Field) SubFields() []string {
	if f.Path == "" {
		return nil
	}
	return strings.Split(f.Path, ".")
}
