package util

type StringSet map[string]empty
type empty struct{}

func NewStringSet(items ...string) StringSet {
	ss := StringSet{}
	ss.Insert(items...)
	return ss
}
func (s StringSet) Insert(items ...string) {
	for _, item := range items {
		s[item] = empty{}
	}
}
