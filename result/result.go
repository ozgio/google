package result

import (
	"github.com/PuerkitoBio/goquery"
)

type Constructor func() Result

// Result
type Result interface {
	Contenter
	Parser
	Name() string
	Priority() int
	Found() bool
}

// Texter
type Contenter interface {
	Text() string
	SimpleHTML() string
	DOM() *goquery.Selection
}

// Parser
type Parser interface {
	Parse(doc *goquery.Document) error
}

type resultBase struct {
	dom        *goquery.Selection
	text       string
	simpleHTML string
}

func (r *resultBase) Text() string {
	return r.text
}

func (r *resultBase) SimpleHTML() string {
	return r.simpleHTML
}

func (r *resultBase) DOM() *goquery.Selection {
	return r.dom
}
