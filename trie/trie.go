package trie

//PRO TIP: if you are having troubles and want to see
//what your trie structure looks like at various points,
//either use the debugger, or try this package:
//https://github.com/davecgh/go-spew

//Trie implements a trie data structure mapping strings to int64s
//that is safe for concurrent use.
type Trie struct {
}

//NewTrie constructs a new Trie.
func NewTrie() *Trie {
	return &Trie{}
}

//Len returns the number of entries in the trie.
func (t *Trie) Len() int {
	panic("implement this function according to the comments above")
}

//Add adds a key and value to the trie.
func (t *Trie) Add(key string, value int64) {
	panic("implement this function according to the comments above")
}

//Find finds `max` values matching `prefix`. If the trie
//is entirely empty, or the prefix is empty, or max == 0,
//or the prefix is not found, this returns a nil slice.
func (t *Trie) Find(prefix string, max int) []int64 {
	panic("implement this function according to the comments above")
}

//Remove removes a key/value pair from the trie
//and trims branches with no values.
func (t *Trie) Remove(key string, value int64) {
	panic("implement this function according to the comments above")
}
