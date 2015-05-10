package fetcher

import (
	"log"
	"net/http"
	"regexp"
	"time"
)

type Fetcher struct {
	Timeout     time.Duration
	Error       error
	AcceptTypes []string
}

func httpClient(f *Fetcher) (p *http.Client) {
	return &http.Client{
		Timeout: f.Timeout,
	}
}

func isAcceptableType(acceptables []string, types []string) bool {
	for _, v := range types {
		for _, acceptable := range acceptables {
			m, _ := regexp.MatchString("^"+acceptable, v)
			if m {
				return true
			}
		}
	}
	return false
}

func (f *Fetcher) checkContentType(url string) bool {
	if len(f.AcceptTypes) == 0 {
		return true
	}

	res := f.Head(url)
	if res == nil {
		return false
	}

	contentTypes := res.Header["Content-Type"]
	result := isAcceptableType(f.AcceptTypes, contentTypes)
	if result == false {
		log.Printf("Expect %s, but %s", f.AcceptTypes, contentTypes)
	}
	return result
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		Timeout:     time.Duration(2 * time.Second),
		AcceptTypes: []string{},
	}
}

func (f *Fetcher) Fetch(url string) *http.Response {
	if f.Error != nil {
		log.Println("error exists: %v", f.Error)
		return nil
	}

	// check Content-Type of the content.
	if len(f.AcceptTypes) > 0 {
		if f.checkContentType(url) == false {
			// TODO: set custom error to f.Error
			return nil
		}
	}

	res, err := httpClient(f).Get(url)
	f.Error = err

	if err != nil {
		log.Fatal("HTTP Get Failed: %v", err)
		return nil
	} else {
		return res
	}
}

func (f *Fetcher) Head(url string) *http.Response {
	if f.Error != nil {
		log.Println("error exists: %v", f.Error)
		return nil
	}

	res, err := httpClient(f).Head(url)
	f.Error = err

	if err != nil {
		log.Fatal("HTTP Head Failed: %v", err)
		return nil
	} else {
		return res
	}
}
