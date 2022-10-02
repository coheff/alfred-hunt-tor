package sources

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/coheff/alfred-hunt-tor/internal/http"
	"github.com/coheff/alfred-hunt-tor/internal/query"
)

// Magnet returns a magnet link for a given url.
func Magnet(url string) string {
	d := http.Document(url)
	r, e := d.
		Find(".torrent-detail-page li a").
		Eq(0).
		Attr("href")
	if !e {
		log.Fatal("Could not find link for: ", url)
	}
	return r
}

// Items returns a slice of scriptfilter.Items for a given query,
// optional category tag, and optional sort flag.
func (lt Leet) Items(q *query.Query) []*Item {
	// both category & sort flag exist
	if q.Cat != "" && q.Sort != "" {
		return scrape(lt.base, fmt.Sprintf("/sort-category-search/%s/%s/%s/desc/1/", q.Query, lt.cats[q.Cat], lt.sorts[q.Sort]))
	}

	// only category exists
	if q.Cat != "" {
		return scrape(lt.base, fmt.Sprintf("/category-search/%s/%s/1/", q.Query, lt.cats[q.Cat]))
	}

	// only sort flag exists
	if q.Sort != "" {
		return scrape(lt.base, fmt.Sprintf("/sort-search/%s/%s/desc/1/", q.Query, lt.sorts[q.Sort]))
	}

	// only query exists
	return scrape(lt.base, fmt.Sprintf("/search/%s/1/", q.Query))
}

// scrape parses a goquery.Document for metadata for a given url base + resource,
// transforming each result and returning as slice of scriptfilter.Items.
func scrape(base, resource string) []*Item {
	var items []*Item

	d := http.Document(base + resource)
	d.
		Find("tbody tr").
		Each(func(i int, s *goquery.Selection) {
			title := s.Find("a").Eq(1).Text()
			seed := s.Find("td").Eq(1).Text()
			time := s.Find("td").Eq(3).Text()
			size := size(s.Find("td").Eq(4).Text())
			uploader := s.Find("td").Eq(5).Text()
			subtitle := seed + " | " + time + " | " + size + " | " + uploader

			r, e := s.
				Find("a").
				Eq(1).
				Attr("href")
			if !e {
				log.Fatal("Could not find link for: ", title)
			}

			items = append(
				items,
				&Item{
					Title:    title,
					Subtitle: subtitle,
					Link:     base + r,
					IconPath: "icons/1337x.png",
				},
			)
		})
	return items
}

// size parses and returns file size from a given string.
func size(s string) string {
	mb := "MB"
	gb := "GB"

	if strings.Contains(s, mb) {
		return strings.SplitAfter(s, mb)[0]
	} else if strings.Contains(s, gb) {
		return strings.SplitAfter(s, gb)[0]
	} else {
		return s
	}
}
