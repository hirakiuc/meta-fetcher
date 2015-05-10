package fetcher

import (
	"encoding/json"
	"net/http"
	"testing"
)

type result struct {
	Args    string
	Headers map[string]string
	Origin  string
	Url     string
}

func parseResult(response *http.Response) *result {
	d := new(result)
	json.NewDecoder(response.Body).Decode(&d)
	return d
}

func TestNewFetcher(t *testing.T) {
	f := NewFetcher()
	if f == nil {
		t.Errorf("NewFetcher return nil.")
	}
}

func TestFetch(t *testing.T) {
	f := NewFetcher()

	const Url string = "http://httpbin.org/get"

	res := f.Fetch(Url)
	defer res.Body.Close()
	if res == nil {
		t.Errorf("fetch failed: %v", f.Error)
	}

	r := parseResult(res)
	if r.Url != Url {
		t.Errorf("Invalid result: %v", r.Url)
	}
}

func TestAcceptTypeFetchFailed(t *testing.T) {
	f := NewFetcher()
	f.AcceptTypes = []string{"application/rdf+xml"}

	const Url string = "http://httpbin.org/get"

	res := f.Fetch(Url)
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()

	if res != nil {
		t.Errorf("expected Content-Type is not %s", res.Header["Content-Type"])
	}
}

func TestAcceptTypeFetch(t *testing.T) {
	f := NewFetcher()
	f.AcceptTypes = []string{"application/json"}

	const Url string = "http://httpbin.org/get"

	res := f.Fetch(Url)
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()

	if res == nil {
		t.Errorf("fetch failed: unexpected Content-Type")
		return
	}

	r := parseResult(res)
	if r.Url != Url {
		t.Errorf("Invalid result: %v", r.Url)
	}
}

func TestRedirectedFetch(t *testing.T) {
	f := NewFetcher()

	res := f.Fetch("http://httpbin.org/redirect/1")
	defer res.Body.Close()
	if res == nil {
		t.Errorf("fetch failed: %v", f.Error)
	}

	r := parseResult(res)
	if r.Url != "http://httpbin.org/get" {
		t.Errorf("Invalid result: %v", r.Url)
	}
}

func TestHead(t *testing.T) {
	f := NewFetcher()

	res := f.Head("http://httpbin.org/get")
	defer res.Body.Close()
	if res == nil {
		t.Errorf("fetch failed: %v", f.Error)
	}

	if res.Status != "200 OK" {
		t.Errorf("fetch failed: %v", res.Status)
	}
}
