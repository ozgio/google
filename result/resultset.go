package result

import "sync"

//TODO convert result slice to map

// ResultSet ..
type ResultSet struct {
	sync.Mutex
	results []Result
}

func (s *ResultSet) Add(r Result) {
	s.Lock()
	defer s.Unlock()

	var i int
	for ; i < len(s.results) && r.Priority() < s.results[i].Priority(); i++ {
	}
	s.results = append(s.results, nil)
	copy(s.results[i+1:], s.results[i:])
	s.results[i] = r
}

func (s *ResultSet) GetByName(n string) Result {
	s.Lock()
	defer s.Unlock()
	for _, r := range s.results {
		if r.Name() == n {
			return r
		}
	}
	return nil
}

func (s *ResultSet) GetAll() []Result {
	s.Lock()
	defer s.Unlock()
	nr := make([]Result, len(s.results))
	copy(nr, s.results)
	return nr
}

func (s *ResultSet) GetTopSpecialResult() Result {
	s.Lock()
	defer s.Unlock()
	for _, r := range s.results {
		if r.Name() != OrganicName {
			return r
		}
	}
	return nil
}

func (s *ResultSet) Remove(name string) Result {
	s.Lock()
	defer s.Unlock()
	for i, r := range s.results {
		if r.Name() == name {
			s.results = append(s.results[:i], s.results[i+1:]...)
			return r
		}
	}
	return nil
}

func (s *ResultSet) Length() int {
	return len(s.results)
}
