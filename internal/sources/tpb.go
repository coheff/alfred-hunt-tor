package sources

import (
	"fmt"

	"github.com/coheff/alfred-hunt-tor/internal/query"
)

type TPB struct {
	Base  string
	Cats  map[string]string
	Sorts map[string]string
}

// Items returns a slice of scriptfilter.Items for a given query,
// optional category tag, and optional sort flag.
func (tpb *TPB) Items(q *query.Query) []*Item {
	// both category & sort flag exist
	if q.Cat != "" && q.Sort != "" {
		return scrape(tpb.Base, fmt.Sprintf("/sort-category-search/%s/%s/%s/desc/1/", q, tpb.Cats[q.Cat], tpb.Sorts[q.Sort]))
	}

	// only category exists
	if q.Cat != "" {
		return scrape(tpb.Base, fmt.Sprintf("/category-search/%s/%s/1/", q, tpb.Cats[q.Cat]))
	}

	// only sort flag exists
	if q.Sort != "" {
		return scrape(tpb.Base, fmt.Sprintf("/sort-search/%s/%s/desc/1/", q, tpb.Sorts[q.Sort]))
	}

	// only query exists
	return scrape(tpb.Base, fmt.Sprintf("/search/%s/1/", q))
}
