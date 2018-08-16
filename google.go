package google

import (
	"context"
	"net/url"
	"strings"
)

const defaultLinkTemplate = "https://www.google.com/search?q=__QUERY__"
const defaultUserAgent = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"

var singleton *Google

func init() {
	singleton = NewGoogle()
}

type Google struct {
	LinkTemplate string
	UserAgent    string
}

func NewGoogle() *Google {
	return &Google{
		LinkTemplate: defaultLinkTemplate,
		UserAgent:    defaultUserAgent,
	}
}

func (g *Google) Search(ctx context.Context, query string) (*Response, error) {
	page, err := g.SendQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	resp := &Response{
		Raw: page,
	}

	err = resp.Parse()
	return resp, err
}

func (g *Google) GetSearchURL(query string) string {
	return strings.Replace(g.LinkTemplate, "__QUERY__", url.QueryEscape(query), -1)
}

func (g *Google) SendQuery(ctx context.Context, query string) ([]byte, error) {
	url := g.GetSearchURL(query)
	headers := map[string]string{
		"User-Agent": g.UserAgent,
	}
	return HTTPRequest(ctx, url, headers)
}

func Search(query string) (*Response, error) {
	ctx := context.Background()

	return singleton.Search(ctx, query)
}
