package google

import (
	"bytes"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/PuerkitoBio/goquery"
	"github.com/ozgio/google/result"
)

// Definition       string
// DefinitionSource string
// Heading          string
// AbstractText     string
// Abstract         string
// AbstractSource   string
// Image            string
// Type             string
// AnswerType       string
// Redirect         string
// DefinitionURL    string
// Answer           string
// AbstractURL      string
// Results          Results
// RelatedTopics    RelatedTopics

// Response represents google search response and its parsed contents
type Response struct {
	URL     string
	Results result.ResultSet
	Raw     []byte
}

var ResultConstructors []result.Constructor

func init() {
	ResultConstructors = []result.Constructor{
		result.NewOrganic,
	}
}

//TODO add context
func (r *Response) Parse() error {
	var err error
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Raw))
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, newResult := range ResultConstructors {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := newResult()
			err = res.Parse(doc)
			if err != nil {
				log.Debugf("parser error in %s: %s")
				return
			}
			if res.Found() {
				r.Results.Add(res)
			}
		}()
	}

	wg.Wait()
	return nil
}

func (r *Response) ResultFound() bool {
	return r.Results.Length() > 0
}

func (r *Response) Organic() *result.Organic {
	res := r.Results.GetByName(result.OrganicName)
	if res == nil {
		return nil
	}
	if org, ok := res.(*result.Organic); ok {
		return org
	}

	return nil
}
