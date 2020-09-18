package trie

import (
	"fmt"
	"sort"
	"strings"
)

// Add will insert a new TrieObject in the Trie
func (t *Node) Add(id, name string, remove ...string) {
	for position, word := range strings.Fields(removeStringList(name, remove...)) {
		cleanedString := cleanString(word)
		if len(cleanedString) < minWordSize {
			continue
		}
		node := t
		for i, runeValue := range cleanedString {
			if _, ok := node.children[runeValue]; !ok {
				child := &Node{
					currentWord:   cleanedString[:i+1],
					children:      make(map[rune]*Node),
					correctWords:  make(map[string]int),
					possibleWords: make(map[string]int),
					possibleData:  make(map[string]*internalOrderData),
					correctData:   make(map[string]*internalOrderData),
				}
				node.children[runeValue] = child
			}
			node = node.children[runeValue]
			if len(cleanedString[i+1:]) == 0 {
				node.isWord = true
				if _, ok := node.correctData[id]; ok {
					data := node.correctData[id]
					data.position = append(data.position, position)
					node.correctData[id] = data
				} else {
					node.correctData[id] = &internalOrderData{id: id, name: name, position: []int{position}}
				}
				node.correctWords[word]++
			} else {
				if _, ok := node.possibleData[id]; ok {
					data := node.possibleData[id]
					data.position = append(data.position, position)
					node.possibleData[id] = data
				} else {
					node.possibleData[id] = &internalOrderData{id: id, name: name, position: []int{position}}
				}
				node.possibleWords[word]++
			}
		}
	}
}

// IsFilled return a boolean value if the the root node has any child
func (t *Node) IsFilled() bool {
	return len(t.children) > 0
}

// HasWord return a boolean value if the word is recorded in the trie
func (t *Node) HasWord(word string) bool {
	node := t
	for _, runeValue := range cleanString(word) {
		if _, ok := node.children[runeValue]; !ok {
			return false
		}
		node = node.children[runeValue]
	}
	return node.isWord
}

// GetPossibleWords return the possible words for the word parameter
func (t *Node) GetPossibleWords(word string) []string {
	node := t
	for _, runeValue := range cleanString(word) {
		if _, ok := node.children[runeValue]; !ok {
			return nil
		}
		node = node.children[runeValue]
	}
	return getKeyListOrderedFromMap(node.possibleWords)
}

// GetCorrectWords return the matching words for the word parameter
func (t *Node) GetCorrectWords(word string) []string {
	node := t
	for _, runeValue := range cleanString(word) {
		if _, ok := node.children[runeValue]; !ok {
			return nil
		}
		node = node.children[runeValue]
	}
	return getKeyListOrderedFromMap(node.correctWords)
}

// GetCorrectIDs return the matching IDs for the word parameter
func (t *Node) GetCorrectIDs(word string) []string {
	node := t
	for _, runeValue := range cleanString(word) {
		if _, ok := node.children[runeValue]; !ok {
			return nil
		}
		node = node.children[runeValue]
	}
	return getKeyListFromObjMap(node.correctData)
}

// GetPossibleIDs return the matching IDs for the word parameter
func (t *Node) GetPossibleIDs(word string) []string {
	node := t
	for _, runeValue := range cleanString(word) {
		if _, ok := node.children[runeValue]; !ok {
			return nil
		}
		node = node.children[runeValue]
	}
	return getKeyListFromObjMap(node.possibleData)
}

// PrintWordData will print the data of a node
func (t *Node) PrintWordData(word string) {
	node := t
	for _, runeValue := range cleanString(word) {
		if _, ok := node.children[runeValue]; !ok {
			return
		}
		node = node.children[runeValue]
	}
	fmt.Println(node.correctData)
	return
}

// SearchByRelevance return the matching IDs for the word parameter ordered by the complete name data and the distance of the searched data
func (t *Node) SearchByRelevance(phrase string) []SearchData {
	var nodes []*Node
	for _, word := range strings.Fields(phrase) {
		cleanedString := cleanString(word)
		if len(cleanedString) < minWordSize {
			continue
		}
		node := t
		for _, runeValue := range cleanedString {
			if _, ok := node.children[runeValue]; !ok {
				break
			}
			node = node.children[runeValue]
		}
		nodes = append(nodes, node)
	}
	return orderMapByRelevance(intersectNodes(nodes))
}

// SearchByRelevancePaginated return the matching IDs for the word parameter ordered by the complete name data and paginates the result
func (t *Node) SearchByRelevancePaginated(phrase string, pagination Pagination) ([]SearchData, Pagination) {
	return paginateList(t.SearchByRelevance(phrase), pagination)
}

func intersectNodes(nodes []*Node) map[string]*internalOrderData {
	var finalKeys map[string]*internalOrderData
	for _, node := range nodes {
		data := node.correctData
		if len(data) == 0 {
			data = node.possibleData
		}
		if finalKeys == nil {
			finalKeys = data
			continue
		}
		intermediateKeys := make(map[string]*internalOrderData)
		for key, values := range data {
			if value, ok := finalKeys[key]; ok && value != nil {
				intSlice := sort.IntSlice(append(finalKeys[key].position, values.position...))
				sort.Sort(intSlice)
				values.position = intSlice
				intermediateKeys[key] = values
			}
		}
		finalKeys = intermediateKeys
	}
	return finalKeys
}

// GetMaximumSizeOfPossibleIds returns the maximum size of ids for the possible words in a node
func (t *Node) GetMaximumSizeOfPossibleIds() int {
	return getMaximumSizeOfPossibleIds(0, t)
}

func getMaximumSizeOfPossibleIds(max int, t *Node) int {
	if len(t.children) == 0 {
		return len(t.possibleData)
	}
	for _, child := range t.children {
		if new := getMaximumSizeOfPossibleIds(max, child); new > max {
			max = new
		}
	}
	if len(t.possibleData) > max {
		max = len(t.possibleData)
	}
	return max
}

// GetMaximumSizeOfCorrectIds returns the maximum size of ids for the correct words in a node
func (t *Node) GetMaximumSizeOfCorrectIds() int {
	return getMaximumSizeOfCorrectIds(0, t)
}

func getMaximumSizeOfCorrectIds(max int, t *Node) int {
	if len(t.children) == 0 {
		return len(t.correctData)
	}
	for _, child := range t.children {
		if new := getMaximumSizeOfCorrectIds(max, child); new > max {
			max = new
		}
	}
	if len(t.correctData) > max {
		println(t.currentWord, len(t.correctData))
		max = len(t.correctData)
	}
	return max
}

// PrintPathToWord returns the maximum size of ids for the correct words in a node
func (t *Node) PrintPathToWord(word string) {
	node := t
	fmt.Println("Printing word:", word)
	cleanedString := cleanString(word)
	for _, runeValue := range cleanedString {
		node = node.children[runeValue]
		fmt.Println(string(runeValue), node.currentWord, node.possibleWords)
	}
}
