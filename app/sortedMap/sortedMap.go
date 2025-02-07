package main

import (
	"fmt"
	"sort"
)

/*
	This is the example for the sorted Map ,
	which is not available by default in golang
*/

type SortedMap[K comparable, V any] struct {
	m    map[K]V
	keys []K
}

func NewSortedMap[K comparable, V any]() *SortedMap[K, V] {
	return &SortedMap[K, V]{
		m:    make(map[K]V),
		keys: make([]K, 0),
	}
}

func (sm *SortedMap[K, V]) Put(key K, value V) {
	if _, exists := sm.m[key]; !exists {
		sm.keys = append(sm.keys, key)
		sort.Slice(sm.keys, func(i, j int) bool {
			return less(sm.keys[i], sm.keys[j])
		})
	}
}

func (sm *SortedMap[K, V]) Len() int {
	return len(sm.m)
}

func (sm *SortedMap[K, V]) Keys() []K {
	return sm.keys
}

func (sm *SortedMap[K, V]) Get(key K) (V, bool) {
	value, exists := sm.m[key]
	return value, exists
}

func (sm *SortedMap[K, V]) Range(f func(key K, value V) bool) {
	for _, k := range sm.keys {
		if !f(k, sm.m[k]) {
			break
		}
	}
}

func less[T comparable](a, b T) bool {
	switch a := any(a).(type) {
	case int:
		return a < any(b).(int)
	}
	return false
}

func main() {
	fmt.Println("This is the example for the user defined sortedMap")
	ordered_map := NewSortedMap[int, string]()
	ordered_map.Put(1, "one")
	ordered_map.Put(2, "two")

	ordered_map.Range(func(key int, value string) bool {
		if key%2 == 0 {
			fmt.Println(key, value)
			return false
		}
		return true
	})
}
