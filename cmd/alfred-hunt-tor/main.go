package main

import (
	"os/exec"
	"strings"

	"github.com/coheff/alfred-hunt-tor/internal/query"
	"github.com/coheff/alfred-hunt-tor/internal/sources"
	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(run)
}

func run() {
	sf()
	wf.WarnEmpty("No matching torrents", "Try another search?")
	wf.SendFeedback()
}

// Sf constructs and orders script filter items from various sources for the input
// query string. The query string may contain a category tag and/or sort order flag.
func sf() {
	arg := wf.Args()[0]
	if len(arg) < 3 {
		wf.Warn("At least three characters required to search", "Keep typing...")
		return
	}

	// Workaround for 1337x.to not including magnet links in search results
	if strings.HasPrefix(arg, "http") {
		exec.
			Command("open", sources.Magnet(arg)).
			Run()
		return
	}

	q := query.Parse(arg)
	if q.Error != "" {
		err := strings.Split(q.Error, ":")
		wf.Warn(err[0], err[1])
		return
	}

	for _, s := range sources.All() {
		for _, i := range s.Items(q) {
			icn := &aw.Icon{
				Value: i.IconPath,
			}
			wf.
				NewItem(i.Title).
				Subtitle(i.Subtitle).
				Arg(i.Link).
				Icon(icn).
				Valid(true)
		}
	}
}
