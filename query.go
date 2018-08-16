package google

import (
	"fmt"
	"strings"
)

type Query struct {
	//Terms is the main query
	Terms string

	//Exclude words from search: Example output: -2016 -2018
	Exclude []string

	//Range is to search for a number range. Min:5, Max:10 => 5..10
	Range QueryRange

	//Site is to search for a specific site
	Site string

	//Filetype restricts results to those of a certain filetype. E.g., PDF, DOCX, TXT, PPT, etc.
	Filetype string

	//InTitle finds pages with a certain word (or words) in the title
	InTitle string

	//AllInTitle is similar to “intitle,” but only results containing all of the specified words in the title tag will be returned.
	AllInTitle string

	//InURL finds pages with a certain word (or words) in the URL
	InURL string

	// Similar to “inurl,” but only results containing all of the specified words in the URL will be returned.
	AllInURL string

	// InText finds pages containing a certain word (or words) somewhere in the content
	InText string

	// AllInText is similar to “intext,” but only results containing all of the specified words somewhere on the page will be returned.
	AllInText string

	//Weather finds the weather for a specific location. This is displayed in a weather snippet, but it also returns results from other “weather” websites.
	Weather string

	//Movie finds information about a specific movie. Also finds movie showtimes if the movie is currently showing near you.
	Movie string

	//Loc finds results from a given area.
	Loc string
}

//QueryRange represents min, max values for google range query as in "Min..Max"
type QueryRange struct {
	Min int
	Max int
}

//Build builds a google query
//FIXME add '"' to special inputs while building
func (q *Query) Build() string {
	tokens := []string{q.Terms}
	if len(q.Exclude) > 0 {
		tokens = append(tokens, "-"+strings.Join(q.Exclude, " -"))
	}
	if q.Range.Min != q.Range.Max {
		tokens = append(tokens, fmt.Sprintf("%d..%d", q.Range.Min, q.Range.Max))
	}

	pre := map[string]string{
		"site":       q.Site,
		"filetype":   q.Filetype,
		"intitle":    q.InTitle,
		"allintitle": q.AllInTitle,
		"inurl":      q.InURL,
		"allinurl":   q.AllInURL,
		"intext":     q.InText,
		"allintext":  q.AllInText,
		"weather":    q.Weather,
		"movie":      q.Movie,
		"loc":        q.Loc,
	}

	for k, v := range pre {
		if v != "" {
			tokens = append(tokens, k+":"+v)
		}
	}

	return strings.Join(tokens, " ")
}
