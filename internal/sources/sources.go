package sources

import "github.com/coheff/alfred-hunt-tor/internal/query"

type Source interface {
	Items(q *query.Query) []*Item
}

type Leet struct {
	base  string
	cats  map[string]string
	sorts map[string]string
}

// TODO: pull base from config
var lt = Leet{
	base: "https://1337x.to",
	cats: map[string]string{
		"#movies": "Movies",
		"#music":  "Music",
		"#tv":     "TV",
	},
	sorts: map[string]string{
		"-s": "seeders",
		"-t": "time",
	},
}

// All returns a slice of all concrete types implementing the source interface.
func All() []Source {
	return []Source{lt}
}
