package query

import "testing"

// TestParse asserts a query can be parsed.
func TestParse(t *testing.T) {
	q := Parse("Lord of the Rings")

	if q.Query != "Lord of the Rings" {
		t.Errorf("query == '%s'; wanted 'Lord of the Rings'", q.Query)
	}

	if q.Cat != "" {
		t.Errorf("cat == '%s'; wanted ''", q.Cat)
	}

	if q.Sort != "" {
		t.Errorf("sort == '%s'; wanted ''", q.Sort)
	}
}

// TestParseCat asserts that a category tag can be parsed from a query.
func TestParseCat(t *testing.T) {
	q := Parse("Lord of the Rings #movies")

	if q.Query != "Lord of the Rings" {
		t.Errorf("query == '%s'; wanted 'Lord of the Rings'", q.Query)
	}

	if q.Cat != "#movies" {
		t.Errorf("cat == '%s'; wanted '#movies'", q.Cat)
	}
}

// TestParseSort asserts that a sort flag can be parsed from a query.
func TestParseSort(t *testing.T) {
	q := Parse("Lord of the Rings -t")

	if q.Query != "Lord of the Rings" {
		t.Errorf("query == '%s'; wanted 'Lord of the Rings'", q.Query)
	}

	if q.Sort != "-t" {
		t.Errorf("sort == '%s'; wanted '-t'", q.Sort)
	}
}

// TestParseSortCat asserts that a both a category tag
// and sort flag can be parsed from a query.
func TestParseSortCat(t *testing.T) {
	q := Parse("Lord of the Rings #movies -t")

	if q.Query != "Lord of the Rings" {
		t.Errorf("query == '%s'; wanted 'Lord of the Rings'", q.Query)
	}

	if q.Cat != "#movies" {
		t.Errorf("cat == '%s'; wanted '#movies'", q.Cat)
	}

	if q.Sort != "-t" {
		t.Errorf("sort == '%s'; wanted '-t'", q.Sort)
	}
}
