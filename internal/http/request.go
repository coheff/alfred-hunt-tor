package http

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Document makes a GET request to a given url, returning a goquery.Document.
func Document(s string) *goquery.Document {
	c := &http.Client{
		Timeout: 5 * time.Second,
	}
	r, e := c.Get(s)
	if e != nil {
		log.Fatal("Error making HTTP request: ", e)
	}
	defer r.Body.Close()

	d, e := goquery.NewDocumentFromReader(r.Body)
	if e != nil {
		log.Fatal("Error loading HTTP response body: ", e)
	}
	return d
}
