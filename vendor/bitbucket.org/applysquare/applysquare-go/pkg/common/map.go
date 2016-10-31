package common

import "sort"

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func SortMapByValue(m map[string]int, minScore int, reversed bool) []string {
	p := make(PairList, 0, len(m))
	for k, v := range m {
		if v < minScore {
			continue
		}
		p = append(p, Pair{k, v})
	}
	if reversed {
		sort.Sort(sort.Reverse(p))
	} else {
		sort.Sort(p)
	}
	var result []string
	for _, v := range p {
		result = append(result, v.Key)
	}
	return result
}

// StringSetToSlice return sorted []string in increasing order
func StringSetToSlice(m map[string]bool) []string {
	var result []string
	for k, v := range m {
		if v {
			result = append(result, k)
		}
	}
	sort.Strings(result)
	return result
}

func StringSetRemove(m map[string]bool, l ...string) {
	for _, v := range l {
		delete(m, v)
	}
}

// RemoveDuplicateString remove duplicate string, and return a new slice, but the order string in slice can be changed
func RemoveDuplicateString(list []string) []string {
	seen := map[string]bool{}
	for _, str := range list {
		seen[str] = true
	}
	return StringSetToSlice(seen)
}

// MergeMap merge map2 to map1
func MergeMap(m1, m2 map[string]interface{}) {
	for k, v := range m2 {
		m1[k] = v
	}
}
