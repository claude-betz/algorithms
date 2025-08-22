package autocomplete

import (
	"sort"
)

func SortDict(dict []string) []string {
	// ideally call once on startup and store
	sort.Strings(dict)
	return dict
}

func upperbound(prefix string) string { return prefix + "\U0010FFFF" }

// returns best k words matching prefix
func Autocomplete(prefix string, dict []string, k int) []string {
	sortedDict := SortDict(dict)

	lo := sort.Search(
		len(sortedDict), func(i int) bool {
			return sortedDict[i] >= prefix
		},
	)
	hi := sort.Search(
		len(sortedDict), func(i int) bool {
			return sortedDict[i] >= upperbound(prefix)
		},
	)
	return sortedDict[lo:hi]
}
