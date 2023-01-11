package sortedkeys

import (
	"sort"
)

// "constaints"
type Ordered interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Iterator[K Ordered, V any] struct {
	map1   map[K]V
	slice1 []K
	Key    K
	Value  V
}

func New[K Ordered, V any](map1 map[K]V) *Iterator[K, V] {
	slice1 := make([]K, 0, len(map1))
	for key := range map1 {
		slice1 = append(slice1, key)
	}
	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})
	return &Iterator[K, V]{
		map1:   map1,
		slice1: slice1,
	}
}

func (it *Iterator[K, V]) Range() bool {
	if len(it.slice1) <= 0 {
		return false
	}
	it.Key = it.slice1[0]
	it.Value = it.map1[it.Key]

	it.slice1 = it.slice1[1:]
	return true
}

func (it *Iterator[K, V]) Ascend() bool {
	return it.Range()
}

func (it *Iterator[K, V]) Descend() bool {
	if len(it.slice1) <= 0 {
		return false
	}
	L := len(it.slice1) - 1
	it.Key = it.slice1[L]
	it.Value = it.map1[it.Key]

	it.slice1 = it.slice1[:L]
	return true
}
