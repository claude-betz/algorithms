package autocomplete_test

import (
	"testing"

	"algorithms/pkg/autocomplete"

	"github.com/stretchr/testify/require"
)

func Test_Sort(t *testing.T) {
	tests := []struct {
		dict     []string
		expected []string
	}{
		{
			dict:     []string{"a", "aab", "aaa"},
			expected: []string{"a", "aaa", "aab"},
		},
	}

	for _, tc := range tests {
		res := autocomplete.SortDict(tc.dict)
		require.ElementsMatch(t, res, tc.expected)
	}
}

func Test_Autocomplete(t *testing.T) {
	tests := []struct {
		dict     []string
		prefix   string
		k        int
		expected []string
	}{
		{
			dict:     []string{"action", "ambrosia", "cart", "amber"},
			prefix:   "am",
			k:        3,
			expected: []string{"amber", "ambrosia"},
		},
	}

	for _, tc := range tests {
		res := autocomplete.Autocomplete(tc.prefix, tc.dict, tc.k)
		require.ElementsMatch(t, res, tc.expected)
	}
}
