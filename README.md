Google Search Library
=====================

google is a go library for querying google and parsing the result page


## Basic Usage

### Simple Search
```go
    resp, _ := google.Search("Metallica site:en.wikipedia.org")
    fmt.Printf("%v", resp.Organic())
```


### Advanced search
```go
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, time.Second*10)
    
    var q google.Query
    q.Terms = "Metallica"
    q.Site = "en.wikipedia.org"

    g := google.NewGoogle()
    resp, _ := g.Search(ctx, q.Build())
    fmt.Printf("%v", resp.Organic())
```


## TODO
- Implement context cancel and timeout in search

