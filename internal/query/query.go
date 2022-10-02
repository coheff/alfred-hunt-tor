package query

import (
	"fmt"
	"strings"
)

type Query struct {
	Query string
	Cat   string
	Sort  string
	Error string
}

const (
	catToken  = "#"
	sortToken = "-"
)

// Parse splits a query string into a search query, category, and sort flag.
func Parse(s string) *Query {
	i := strings.Index(s, catToken)
	j := strings.Index(s, sortToken)

	// both category & sort flag exist
	if i != -1 && j != -1 {
		c, e1 := parseFlag(s, i)
		q := strings.Replace(s, " "+c, "", -1)

		sf, e2 := parseFlag(s, j)
		q = strings.Replace(q, " "+sf, "", -1)

		if e1 != "" {
			return &Query{q, c, sf, e1}
		}
		return &Query{q, c, sf, e2}
	}

	// only category exists
	if i != -1 {
		c, e := parseFlag(s, i)
		q := strings.Replace(s, " "+c, "", -1)
		return &Query{q, c, "", e}
	}

	// only sort flag exists
	if j != -1 {
		sf, e := parseFlag(s, j)
		q := strings.Replace(s, " "+sf, "", -1)
		return &Query{q, "", sf, e}
	}

	// only query exists
	return &Query{s, "", "", ""}
}

// parseFlag parses a flag (category tag or sort flag) from a given string
// and flag start index.
// A flag can be found either within or at the end of a query string.
func parseFlag(s string, i int) (flag, err string) {
	var f string
	if strings.Contains(s[i:], " ") {
		f = s[i : strings.Index(s[i:], " ")+i]
	} else {
		f = s[i:]
	}

	e := validate(f)
	return f, e
}

// Validate returns an error string if flag isn't recognised, else nil.
func validate(flag string) (err string) {
	cats := []string{"#movies", "#music", "#tv"}
	sorts := []string{"-s", "-t"}
	for _, s := range append(cats, sorts...) {
		if s == flag {
			return ""
		}
	}

	if strings.HasPrefix(flag, catToken) {
		return fmt.Sprintf("Category not recognised:Try one of %s", strings.Join(cats[:], ", "))
	}
	return fmt.Sprintf("Sort flag not recognised:Try one of %s", strings.Join(sorts[:], ", "))
}
