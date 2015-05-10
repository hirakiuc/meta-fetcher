package fetcher

import (
	"log"
	"net/http"
)

func NewHtmlFetcher() *Fetcher {
	f := NewFetcher()
	f.AcceptTypes = []string{"text/html"}
	return f
}

func (f *Fetcher) FetchHtml(url string) *http.Response {
	if f.Error != nil {
		log.Println("error exists: %v", f.Error)
		return nil
	}

	return f.Fetch(url)
}
