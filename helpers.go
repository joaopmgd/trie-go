package trie

import (
	"math"
	"sort"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func getKeyListFromMap(m map[string]int) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func getKeyListFromObjMap(m map[string]*internalOrderData) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func orderMapByRelevance(m map[string]*internalOrderData) []SearchData {
	values := make([]*internalOrderData, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	sort.Sort(byRelevance(values))
	var orderedSearchData []SearchData
	for _, value := range values {
		orderedSearchData = append(orderedSearchData, SearchData{ID: value.id, Name: value.name})
	}
	return orderedSearchData
}

func getKeyListOrderedFromMap(m map[string]int) []string {
	type kv struct {
		key   string
		value int
	}
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].value > ss[j].value
	})
	var orderedSlice []string
	for _, kv := range ss {
		orderedSlice = append(orderedSlice, kv.key)
	}
	return orderedSlice
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func removeStringList(value string, removeStrings ...string) string {
	for _, remove := range removeStrings {
		value = strings.ReplaceAll(value, remove, " ")
	}
	return value
}

func cleanString(value string) string {
	lowerCase := strings.ToLower(value)
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	newValue, _, _ := transform.String(t, lowerCase)
	return rxp.ReplaceAllString(newValue, "")
}

func paginateList(list []SearchData, pagination Pagination) ([]SearchData, Pagination) {
	if len(list) == 0 || pagination.Offset() >= len(list) {
		return nil, pagination
	}
	initialItem := len(list) - 1
	if pagination.Offset() < initialItem {
		initialItem = pagination.Offset()
	}
	finalItem := len(list)
	if pagination.Offset()+int(pagination.PerPage) < finalItem {
		finalItem = pagination.Offset() + int(pagination.PerPage)
	}
	pagination.Total = int32(len(list))
	return list[initialItem:finalItem], pagination
}

// Offset returns the page from the request data
func (p Pagination) Offset() int {
	page := int(math.Max(1, float64(p.Page)))
	return (page - 1) * int(p.PerPage)
}

func dist(i, j []int) int {
	for len(i) > 0 && len(j) > 0 {
		if i[0] < j[0] {
			return -1
		} else if i[0] > j[0] {
			return 1
		}
		i = i[1:]
		j = j[1:]
	}
	if len(i) == len(j) {
		return 0
	} else if len(i) < len(j) {
		return -1
	}
	return 1
}
