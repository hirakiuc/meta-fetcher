package fetcher

import (
	"testing"
)

func TestNewHtmlFetcher(t *testing.T) {
	f := NewHtmlFetcher()
	if f == nil {
		t.Errorf("NewHtmlFetcher return nil.")
	}
}

func TestFetchHtmlUnAcceptedType(t *testing.T) {
	f := NewHtmlFetcher()
	if f == nil {
		t.Errorf("NewHtmlFetcher return nil")
	}

	const Url string = "http://httpbin.org/response-headers?Content-Type=application/json"
	res := f.FetchHtml(Url)
	if res != nil {
		t.Errorf("Unexpected Response %v", res)
	}
}

func TestFetchHtmlAcceptedType(t *testing.T) {
	f := NewHtmlFetcher()
	if f == nil {
		t.Errorf("NewRdfFetcher return nil")
	}
	const Url string = "http://httpbin.org/response-headers?Content-Type=text/html"
	res := f.FetchHtml(Url)
	if res == nil {
		t.Errorf("Html Fetch Failed")
	}
}
