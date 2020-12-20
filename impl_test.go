package trie

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_AddPhrases(t *testing.T) {
	cases := map[string]struct {
		id                    string
		name                  string
		expectedValue         string
		expectedCorrectNames  []string
		expectedCorrectIDs    []string
		possibleValue         string
		expectedPossibleNames []string
		expectedPossibleIDs   []string
	}{
		"Adding one word":          {"1", "direito", "direito", []string{"direito"}, []string{"1"}, "direit", []string{"direito"}, []string{"1"}},
		"Adding two words":         {"1", "direito penal", "direito", []string{"direito"}, []string{"1"}, "direit", []string{"direito"}, []string{"1"}},
		"Adding multiple words":    {"1", "direito penal judiciário do nordeste", "judiciario", []string{"judiciário"}, []string{"1"}, "judiciar", []string{"judiciário"}, []string{"1"}},
		"Adding special character": {"1", "Administração", "administracao", []string{"Administração"}, []string{"1"}, "admin", []string{"Administração"}, []string{"1"}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			trie := NewNode()
			trie.Add(tc.id, tc.name)
			diff := cmp.Diff(tc.expectedCorrectNames, trie.GetCorrectWords(tc.expectedValue))
			if diff != "" {
				t.Fatalf(diff)
			}
			diff = cmp.Diff(tc.expectedCorrectIDs, trie.GetCorrectIDs(tc.expectedValue))
			if diff != "" {
				t.Fatalf(diff)
			}
			diff = cmp.Diff(tc.expectedPossibleNames, trie.GetPossibleWords(tc.possibleValue))
			if diff != "" {
				t.Fatalf(diff)
			}
			diff = cmp.Diff(tc.expectedPossibleIDs, trie.GetPossibleIDs(tc.possibleValue))
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

// func Test_AddWordsSequentially(t *testing.T) {
// 	trie := NewNode()
// 	cases := []struct {
// 		testName              string
// 		id                    string
// 		name                  string
// 		expectedValue         string
// 		expectedCorrectNames  []string
// 		expectedCorrectIDs    []string
// 		possibleValue         string
// 		expectedPossibleNames []string
// 		expectedPossibleIDs   []string
// 	}{
// 		{"Adding one word", "1", "direito", "direito", []string{"direito"}, []string{"1"}, "direit", []string{"direito"}, []string{"1"}},
// 		{"Adding two words", "2", "direito penal", "direito", []string{"direito"}, []string{"1", "2"}, "direit", []string{"direito"}, []string{"1", "2"}},
// 		{"Adding multiple words", "3", "direito penal judiciário do nordeste", "judiciario", []string{"judiciário"}, []string{"3"}, "judiciar", []string{"judiciário"}, []string{"3"}},
// 		{"Adding special character", "4", "Administração", "administracao", []string{"Administração"}, []string{"4"}, "admin", []string{"Administração"}, []string{"4"}},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.testName, func(t *testing.T) {
// 			trie.Add(tc.id, tc.name)
// 			diff := cmp.Diff(tc.expectedCorrectNames, trie.GetCorrectWords(tc.expectedValue))
// 			if diff != "" {
// 				t.Fatalf(diff)
// 			}
// 			diff = cmp.Diff(tc.expectedCorrectIDs, trie.GetCorrectIDs(tc.expectedValue))
// 			if diff != "" {
// 				t.Fatalf(diff)
// 			}
// 			diff = cmp.Diff(tc.expectedPossibleNames, trie.GetPossibleWords(tc.possibleValue))
// 			if diff != "" {
// 				t.Fatalf(diff)
// 			}
// 			diff = cmp.Diff(tc.expectedPossibleIDs, trie.GetPossibleIDs(tc.possibleValue))
// 			if diff != "" {
// 				t.Fatalf(diff)
// 			}
// 		})
// 	}
// }

func Test_IsFilled(t *testing.T) {
	cases := map[string]struct {
		id       string
		name     string
		isFilled bool
	}{
		"Is filled":                      {"1", "direito penal", true},
		"Is NOT filled":                  {"1", "", false},
		"Is NOT filled with small words": {"1", "oi", false},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			trie := NewNode()
			trie.Add(tc.id, tc.name)
			if trie.IsFilled() != tc.isFilled {
				t.Fatalf("\nExpected: %v\nGot: %v", tc.isFilled, trie.IsFilled())
			}
		})
	}
}

func Test_HasWord(t *testing.T) {
	cases := map[string]struct {
		id      string
		name    string
		word    string
		hasWord bool
	}{
		"Has the word":          {"1", "direito", "direito", true},
		"Dos not have the word": {"1", "direito", "penal", false},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			trie := NewNode()
			trie.Add(tc.id, tc.name)
			if trie.HasWord(tc.word) != tc.hasWord {
				t.Fatalf("\nExpected: %v\nGot: %v", tc.hasWord, trie.HasWord(tc.word))
			}
		})
	}
}

func Test_GetMaximumSizeOfPossibleIds(t *testing.T) {
	cases := map[string]struct {
		data []struct {
			id   string
			name string
		}
		maxPossible int
	}{
		"Maximum size is 8": {[]struct {
			id   string
			name string
		}{{"1", "Administração"}, {"2", "Administracao"}, {"3", "Administraçao"}, {"4", "Administracão"}, {"5", "administração"}, {"6", "administracao"}, {"7", "administraçao"}, {"8", "administracão"}}, 8},
		"Maximum size is 0": {nil, 0},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			trie := NewNode()
			for _, d := range tc.data {
				trie.Add(d.id, d.name)
			}
			if trie.GetMaximumSizeOfPossibleIds() != tc.maxPossible {
				t.Fatalf("\nExpected: %v\nGot: %v", trie.GetMaximumSizeOfPossibleIds(), tc.maxPossible)
			}
		})
	}
}

func Test_GetMaximumSizeOfCorrectIds(t *testing.T) {
	cases := map[string]struct {
		data []struct {
			id   string
			name string
		}
		maxCorrect int
	}{
		"Maximum size is 8": {[]struct {
			id   string
			name string
		}{{"1", "Administração"}, {"2", "Administracao"}, {"3", "Administraçao"}, {"4", "Administracão"}, {"5", "administração"}, {"6", "administracao"}, {"7", "administraçao"}, {"8", "administracão"}}, 8},
		"Maximum size is 0": {nil, 0},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			trie := NewNode()
			for _, d := range tc.data {
				trie.Add(d.id, d.name)
			}
			if trie.GetMaximumSizeOfCorrectIds() != tc.maxCorrect {
				t.Fatalf("\nExpected: length of %v\nGot: length of %v", trie.GetMaximumSizeOfCorrectIds(), tc.maxCorrect)
			}
		})
	}
}

func Test_SearchByRelevance(t *testing.T) {
	trieNode := NewNode()
	cases := []struct {
		testName string
		id       string
		name     string
		word     string
		expected []SearchData
	}{
		{"Adding one word 1", "1", "Direito Penal", "direito penal", []SearchData{{"1", "Direito Penal"}}},
		{"Adding one word 2", "2", "Direito Penal Militar", "direito penal", []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}}},
		{"Adding one word 3", "3", "Direito Penal / Princípios do Direito Penal", "direito penal", []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}}},
		{"Adding one word 4", "4", "Direito Penal / Introdução ao estudo do Direito Penal", "direito penal", []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}, {"4", "Direito Penal / Introdução ao estudo do Direito Penal"}}},
		{"Adding one word 5", "5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito", "direito penal", []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}, {"4", "Direito Penal / Introdução ao estudo do Direito Penal"}, {"5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito"}}},
		{"Adding one word 6", "6", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Grego", "direito penal", []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}, {"4", "Direito Penal / Introdução ao estudo do Direito Penal"}, {"5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito"}, {"6", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Grego"}}},
		{"Adding one word 7", "7", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Romano", "direito penal", []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}, {"4", "Direito Penal / Introdução ao estudo do Direito Penal"}, {"5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito"}, {"6", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Grego"}, {"7", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Romano"}}},
		{"Adding one word 8", "8", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal e o IIuminismo", "direito penal", []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}, {"4", "Direito Penal / Introdução ao estudo do Direito Penal"}, {"5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito"}, {"6", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Grego"}, {"7", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Romano"}, {"8", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal e o IIuminismo"}}},
		{"Adding one word 9", "9", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal na Idade Média", "direito penal", []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}, {"4", "Direito Penal / Introdução ao estudo do Direito Penal"}, {"5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito"}, {"6", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Grego"}, {"7", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Romano"}, {"8", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal e o IIuminismo"}, {"9", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal na Idade Média"}}},
		{"Adding one word 10", "10", "Direito Penal / Introdução ao estudo do Direito Penal / As Velocidades do Direito Penal", "direito penal", []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}, {"4", "Direito Penal / Introdução ao estudo do Direito Penal"}, {"5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito"}, {"6", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Grego"}, {"7", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Romano"}, {"8", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal e o IIuminismo"}, {"9", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal na Idade Média"}, {"10", "Direito Penal / Introdução ao estudo do Direito Penal / As Velocidades do Direito Penal"}}},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			trieNode.Add(tc.id, tc.name)
			diff := cmp.Diff(tc.expected, trieNode.SearchByRelevance(tc.word))
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func Test_SearchByRelevancePaginated(t *testing.T) {
	trieNode := NewNode()
	cases := []struct {
		testName           string
		id                 string
		name               string
		word               string
		pagination         Pagination
		expectedData       []SearchData
		expectedPagination Pagination
	}{
		{"Adding one word 1", "1", "Direito Penal", "direito penal", Pagination{PerPage: 3, Page: 1}, []SearchData{{"1", "Direito Penal"}}, Pagination{PerPage: 3, Page: 1, Total: 1}},
		{"Adding one word 2", "2", "Direito Penal Militar", "direito penal", Pagination{PerPage: 3, Page: 1}, []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}}, Pagination{PerPage: 3, Page: 1, Total: 2}},
		{"Adding one word 3", "3", "Direito Penal / Princípios do Direito Penal", "direito penal", Pagination{PerPage: 3, Page: 1}, []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}}, Pagination{PerPage: 3, Page: 1, Total: 3}},
		{"Adding one word 4", "4", "Direito Penal / Introdução ao estudo do Direito Penal", "direito penal", Pagination{PerPage: 3, Page: 1}, []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}}, Pagination{PerPage: 3, Page: 1, Total: 4}},
		{"Adding one word 5", "5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito", "direito penal", Pagination{PerPage: 3, Page: 1}, []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}}, Pagination{PerPage: 3, Page: 1, Total: 5}},
		{"Adding one word 6", "6", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Grego", "direito penal", Pagination{PerPage: 3, Page: 1}, []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}}, Pagination{PerPage: 3, Page: 1, Total: 6}},
		{"Adding one word 7", "7", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Romano", "direito penal", Pagination{PerPage: 3, Page: 2}, []SearchData{{"4", "Direito Penal / Introdução ao estudo do Direito Penal"}, {"5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito"}, {"6", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Grego"}}, Pagination{PerPage: 3, Page: 2, Total: 7}},
		{"Adding one word 8", "8", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal e o IIuminismo", "direito penal", Pagination{PerPage: 3, Page: 3}, []SearchData{{"7", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Romano"}, {"8", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal e o IIuminismo"}}, Pagination{PerPage: 3, Page: 3, Total: 8}},
		{"Adding one word 9", "9", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal na Idade Média", "direito penal", Pagination{PerPage: 10, Page: 2}, nil, Pagination{PerPage: 10, Page: 2}},
		{"Adding one word 10", "10", "Direito Penal / Introdução ao estudo do Direito Penal / As Velocidades do Direito Penal", "direito penal", Pagination{PerPage: 100, Page: 1}, []SearchData{{"1", "Direito Penal"}, {"2", "Direito Penal Militar"}, {"3", "Direito Penal / Princípios do Direito Penal"}, {"4", "Direito Penal / Introdução ao estudo do Direito Penal"}, {"5", "Direito Penal / Introdução ao estudo do Direito Penal / O Direito Penal e o Estado Democrático de Direito"}, {"6", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Grego"}, {"7", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal Romano"}, {"8", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal e o IIuminismo"}, {"9", "Direito Penal / Introdução ao estudo do Direito Penal / Evolução Histórica / Direito Penal na Idade Média"}, {"10", "Direito Penal / Introdução ao estudo do Direito Penal / As Velocidades do Direito Penal"}}, Pagination{PerPage: 100, Page: 1, Total: 10}},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			trieNode.Add(tc.id, tc.name)
			data, pag := trieNode.SearchByRelevancePaginated(tc.word, tc.pagination)
			diff := cmp.Diff(tc.expectedData, data)
			if diff != "" {
				t.Fatalf(diff)
			}
			diff = cmp.Diff(tc.expectedPagination, pag)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
