package trie

import "regexp"

// Min word size to get in the Trie
const minWordSize = 3

// Regex is declared as global to the package so it is not compiled on every execution
var rxp = regexp.MustCompile("[^A-Za-zÀ-ÖØ-öø-ÿ0-9-_]+")

// Node is the data structure that hold IDs and runes of an object
type Node struct {
	possibleData  map[string]*internalOrderData
	correctData   map[string]*internalOrderData
	possibleWords map[string]int
	correctWords  map[string]int
	currentWord   string
	isWord        bool
	children      map[rune]*Node
}

// Pagination data for selecting the number of ids in the trie
type Pagination struct {
	PerPage int32
	Page    int32
	Total   int32
}

// SearchData returns the ID and name found in the trie
type SearchData struct {
	ID   string
	Name string
}

type internalOrderData struct {
	id       string
	name     string
	position []int
}

type byRelevance []*internalOrderData

func (n byRelevance) Len() int { return len(n) }
func (n byRelevance) Less(i, j int) bool {
	return dist(n[i].position, n[j].position) == -1 ||
		(dist(n[i].position, n[j].position) == 0 && len(n[i].name) < len(n[j].name)) ||
		(dist(n[i].position, n[j].position) == 0 && len(n[i].name) == len(n[j].name) && n[i].name < n[j].name)
}
func (n byRelevance) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
