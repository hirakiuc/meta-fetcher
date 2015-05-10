package fetcher

import (
	"encoding/xml"
	"log"
	"time"
)

const RdfType string = "application/rdf+xml"

// To convert to time.Time, see http://play.golang.org/p/JBi5Qu470_
//
type datetime time.Time

func (t *datetime) UnmarshalText(b []byte) error {
	result, err := time.Parse("2015-01-18T15:03:59+09:00", string(b))
	if err == nil {
		*t = datetime(result)
		return nil
	}

	var t2 time.Time
	err = t2.UnmarshalText(b)
	if err == nil {
		*t = datetime(t2)
		return nil
	}

	return err
}

func (t datetime) Time() time.Time {
	return (time.Time)(t)
}

type FeedItem struct {
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Date          datetime `xml:"http://purl.org/dc/elements/1.1/ date"`
	Subject       string   `xml:"http://purl.org/dc/elements/1.1/ subject"`
	Bookmarkcount int32    `xml:"http://www.hatena.ne.jp/info/xmlns# bookmarkcount"`
}

type Rdf struct {
	Items []FeedItem `xml:"item"`
}

func NewRdfFetcher() *Fetcher {
	f := NewFetcher()
	f.AcceptTypes = []string{"application/rdf+xml", "application/xml"}
	return f
}

func (f *Fetcher) FetchRdf(url string) *Rdf {
	if f.Error != nil {
		log.Println("error exists: %v", f.Error)
		return nil
	}

	res := f.Fetch(url)
	if res == nil {
		return nil
	}
	defer res.Body.Close()

	rdf := &Rdf{Items: []FeedItem{}}
	f.Error = xml.NewDecoder(res.Body).Decode(&rdf)
	if f.Error != nil {
		log.Println("Rdf Parse Failed: %v", f.Error)
		return rdf
	}

	return rdf
}
