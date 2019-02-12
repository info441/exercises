package trie

//int64set is a set of int64 values
type int64set map[int64]struct{}

//add adds a value to the set and returns
//true if the value didn't already exist in the set.
func (s int64set) add(value int64) bool {
	if exist := s.has(value); !exist {
		s[value] = struct{}{}
		return !exist
	}
	return false
}

//remove removes a value from the set and returns
//true if that value was in the set, false otherwise.
func (s int64set) remove(value int64) bool {
	if exist := s.has(value); exist {
		delete(s, value)
		return exist
	}
	return false
}

//has returns true if value is in the set,
//or false if it is not in the set.
func (s int64set) has(value int64) bool {
	_, exist := s[value]
	return exist
}

//all returns all values in the set as a slice.
//The returned slice will always be non-nil, but
//the order will be random. Use sort.Slice to
//sort the slice if necessary.
func (s int64set) all() []int64 {
	keys := make([]int64, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}