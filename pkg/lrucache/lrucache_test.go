package lrucache_test

import (
	"algorithms/pkg/lrucache"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LRUCache(t *testing.T) {
	lru, err := lrucache.NewLRU(3)
	fmt.Printf("new lru: %s\n", lru.String())
	require.NoError(t, err)
	lru.Put(1, 1)
	fmt.Printf("insert 1: %s\n", lru.String())
	lru.Put(2, 3)
	fmt.Printf("insert 2: %s\n", lru.String())
	lru.Put(4, 4)
	fmt.Printf("insert 3: %s\n", lru.String())
	lru.Put(5, 5)
	fmt.Printf("insert 4: %s\n", lru.String())

	_, ok := lru.Get(1)
	require.False(t, ok)
	fmt.Printf("get 1: %s\n", lru.String())
	res, ok := lru.Get(2)
	require.True(t, ok)
	fmt.Printf("get 2: %s\n", lru.String())
	require.Equal(t, 3, res)
}
