package ordered_map

import (
	"fmt"
	"strings"
)

type OrderedMap[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Remove(key K)
	Keys() []K
	Values() []V
	Clear()
	Len() int
	ForEach(fn func(key K, value V) bool)
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

type orderedMap[K comparable, V any] struct {
	entries []entry[K, V]
	index   map[K]int
}

func NewOrderedMap[K comparable, V any]() OrderedMap[K, V] {
	return &orderedMap[K, V]{
		index: make(map[K]int),
	}
}

func (om *orderedMap[K, V]) Set(key K, value V) {
	if idx, ok := om.index[key]; ok {
		om.entries[idx].value = value
		return
	}

	om.index[key] = len(om.entries)
	om.entries = append(om.entries, entry[K, V]{key: key, value: value})
}

func (om *orderedMap[K, V]) Get(key K) (V, bool) {
	if idx, ok := om.index[key]; ok {
		return om.entries[idx].value, true
	}
	var zero V
	return zero, false
}

func (om *orderedMap[K, V]) Remove(key K) {
	idx, ok := om.index[key]
	if !ok {
		return
	}

	delete(om.index, key)

	n := 0
	for i, e := range om.entries {
		if i != idx {
			om.entries[n] = e
			om.index[e.key] = n
			n++
		}
	}

	var zero entry[K, V]
	om.entries[len(om.entries)-1] = zero
	om.entries = om.entries[:n]
}

func (om *orderedMap[K, V]) Keys() []K {
	keys := make([]K, len(om.entries))
	for i, e := range om.entries {
		keys[i] = e.key
	}
	return keys
}

func (om *orderedMap[K, V]) Values() []V {
	values := make([]V, len(om.entries))
	for i, e := range om.entries {
		values[i] = e.value
	}
	return values
}

func (om *orderedMap[K, V]) Clear() {
	om.entries = nil
	om.index = make(map[K]int)
}

func (om *orderedMap[K, V]) Len() int {
	return len(om.entries)
}

func (om *orderedMap[K, V]) ForEach(fn func(key K, value V) bool) {
	for _, e := range om.entries {
		if !fn(e.key, e.value) {
			break
		}
	}
}

func (om *orderedMap[K, V]) String() string {
	var b strings.Builder
	b.WriteString("OrderedMap[")
	for i, e := range om.entries {
		if i > 0 {
			b.WriteString(" ")
		}
		fmt.Fprintf(&b, "{%v: %v}", e.key, e.value)
	}
	b.WriteString("]")
	return b.String()
}
