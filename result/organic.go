package result

import (
	"errors"
	"html"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

const OrganicName = "organic"
const organicID = 100

func NewOrganic() Result {
	return &Organic{}
}

// Organic
type Organic struct {
	resultBase
	links Links
}

func (r *Organic) Priority() int {
	return organicID
}

func (r *Organic) Name() string {
	return OrganicName
}

func (r *Organic) Links() Links {
	return r.links
}

func (r *Organic) Found() bool {
	return len(r.links) > 0
}

func (r *Organic) Parse(doc *goquery.Document) error {
	no := 1
	doc.Find(".g").Each(func(i int, s *goquery.Selection) {
		var html string
		if log.GetLevel() == log.DebugLevel {
			html, _ = s.Html()
		}

		link := s.Find(".r a")
		url, err := getURLInsideATag(link)
		if err != nil {
			log.Debugf("%s, i: %d", err.Error(), i)
			log.Debugln(html)
			return
		}

		abstract := s.Find(".s .st")
		if abstract.Text() == "" {
			log.Debugf("abstract is empty, i: %d", i)
			log.Debugln(html)
			return //TODO ?
		}

		r.links = append(r.links, Link{
			No:       no,
			Title:    link.Text(),
			URL:      url,
			Abstract: abstract.Text(),
		})
		no++

	})
	return nil
}

func getURLInsideATag(link *goquery.Selection) (string, error) {
	if link.Length() == 0 {
		return "", errors.New("link cannot be found")
	}

	url, urlFound := link.Attr("href")
	if !urlFound {
		return "", errors.New("cannot find the href attribute")
	}

	url, err := getTargetLink(url)
	return url, err
}

func getTargetLink(href string) (string, error) {
	href = html.UnescapeString(href)
	u, err := url.Parse(href)
	if err != nil {
		return "", err
	}
	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", err
	}
	link, found := query["q"]
	if !found {
		return "", errors.New("Cannot find 'q' in path")
	}
	return link[0], nil
}
