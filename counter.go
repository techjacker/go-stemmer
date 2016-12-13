package main

// Counter will count the occurrences of text added to it
type Counter interface {
	Add(in []byte)
	GetOccurrences() map[string]int
}

// NewCounter is a Counter factory
func NewCounter() Counter {
	return &counter{
		occurrences: make(map[string]int),
	}
}

type counter struct {
	occurrences map[string]int
}

// need to convert to array to use as key in map
func (s *counter) Add(in []byte) {
	if in != nil {
		k := string(in)
		if _, ok := s.occurrences[k]; ok {
			s.occurrences[k]++
		} else {
			s.occurrences[k] = 1
		}
	}
}

func (s *counter) GetOccurrences() map[string]int {
	return s.occurrences
}
