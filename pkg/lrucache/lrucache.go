package lrucache

import (
	"container/list"
	"fmt"
	"strings"
)

type pair struct {
	key int
	val int
}

func (p pair) String() string {
	return fmt.Sprintf("%d:%d", p.key, p.val)
}

// Least recently used cache
type LRU struct {
	capacity int
	list     *list.List
	index    map[int]*list.Element
}

func NewLRU(capacity int) (*LRU, error) {
	if capacity < 1 {
		return nil, fmt.Errorf("capacity must be at least 1")
	}
	return &LRU{
		capacity: capacity,
		list:     list.New(),
		index:    make(map[int]*list.Element, capacity),
	}, nil
}

func (l *LRU) String() string {
	count := 0
	var b strings.Builder
	b.WriteString("[ ")
	for e := l.list.Front(); e != nil; e = e.Next() {
		pair, ok := e.Value.(pair)
		if !ok {
			return "invalid"
		}
		b.WriteString(fmt.Sprintf("%d[%s] ", count, pair))
		count++
	}
	b.WriteString("]\n")
	return b.String()
}

func (l *LRU) Put(k, val int) {
	toInsert := pair{key: k, val: val}
	e, ok := l.index[k]
	if ok {
		// value exists, update value, move to front of list
		e.Value = toInsert
		l.list.MoveToFront(e)
		return
	} else {
		// value doesn't exist, pushback new value
		e = l.list.PushFront(toInsert)
		// update index
		l.index[k] = e
		// handle size
		if l.list.Len() > l.capacity {
			last := l.list.Back()
			pair, ok := last.Value.(pair)
			if !ok {
				return
			}
			l.list.Remove(last)
			delete(l.index, pair.key)
		}
	}
}

func (l *LRU) Get(k int) (int, bool) {
	e, ok := l.index[k]
	if !ok {
		return -1, false
	}
	l.list.MoveToFront(e)
	pair, ok := e.Value.(pair)
	if !ok {
		return -1, false
	}
	return pair.val, true
}
