package fetcher

import (
	//  "net/http"
	"testing"
)

func TestNewRdfFetcher(t *testing.T) {
	f := NewRdfFetcher()
	if f == nil {
		t.Errorf("NewRdfFetcher return nil.")
	}
}

func TestFetchUnAcceptedType(t *testing.T) {
	f := NewRdfFetcher()
	if f == nil {
		t.Errorf("NewRdfFetcher return nil")
	}

	const Url string = "http://httpbin.org/response-headers?Content-Type=application/json"
	items := f.FetchRdf(Url)
	if items != nil {
		t.Errorf("Unexpected Response %v", items)
	}
}

func TestFetchAcceptedType(t *testing.T) {
	f := NewRdfFetcher()
	if f == nil {
		t.Errorf("NewRdfFetcher return nil")
	}
	const Url string = "http://feeds.feedburner.com/hatena/b/hotentry"
	rdf := f.FetchRdf(Url)
	if rdf == nil {
		t.Errorf("Rdf Fetch Failed")
	}
	if len(rdf.Items) == 0 {
		t.Errorf("Rdf.Items is empty...")
	}
}
