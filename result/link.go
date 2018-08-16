package result

// Link represent a single item in organic search results
type Link struct {
	Title    string
	URL      string
	Abstract string
	No       int
}

type Links []Link

func (links *Links) GetByNo(no int) *Link {
	if no < 1 || no > len(*links) {
		return nil
	}

	for _, res := range *links {
		if res.No == no {
			return &res
		}
	}

	return nil
}
