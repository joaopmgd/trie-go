package trie

// NodeInterface is the interface satisfied by the Trie
type NodeInterface interface {
	// Add a new object to the Trie
	// the remove string list parameter will remove the patterns and transform them in spaces
	// so if the pattern is found in the middle of a word, then it will became two words with the pattern removed
	Add(id, name string, remove ...string)
	// Checks if the Trie has at least one object
	IsFilled() bool
	// Checks if word is in the Trie
	HasWord(word string) bool
	// Based on a word, get the possible words following from that
	GetPossibleWords(word string) []string
	// Based on a word, get the correct matching words that had been inserted
	// if the word is not found than the possible words are appended
	SearchByRelevance(phrase string) []SearchData
	// It is the same as Search by name, but is paginated the final slice is paginated and ordered by its name
	// The pagination returned has the total data and the number of page items
	SearchByRelevancePaginated(phrase string, pagination Pagination) ([]SearchData, Pagination)
}

// NodeHelperInterface is an extra interface that the trie implements
// So it makes simpler to implement the node interface
type NodeHelperInterface interface {
	// Based on a word, get the correct matching words that had been inserted
	GetCorrectWords(word string) []string
	// Based on a word, get the correct matching IDs that had been inserted
	GetCorrectIDs(word string) []string
	// Based on a word, get the possible IDs following from that
	GetPossibleIDs(word string) []string
	// Get the maximum node size of possible IDs
	GetMaximumSizeOfPossibleIds() int
	// Get the maximum node size of correct IDs
	GetMaximumSizeOfCorrectIds() int
	// Based on a word, print data from the root of the trie until it reaches the final rune
	PrintPathToWord(word string)
	// It is the same as Search by name, but is paginated the final slice is paginated and ordered by its name
	// The pagination returned has the total data and the number of page items
	PrintWordData(word string)
}

// NewNode returns a Trie ready to be used
func NewNode() *Node {
	return &Node{children: make(map[rune]*Node)}
}
