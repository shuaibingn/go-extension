package ordered_map

import (
	"container/list"
	"fmt"
)

type OrderedMap[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Remove(key K)
	Keys() []K
	Values() []V
	Clear()
	Len() int
	Iterator() <-chan listValue[K, V]
}

type orderedMap[K comparable, V any] struct {
	m    map[K]*list.Element
	list *list.List
}

func (om *orderedMap[K, V]) String() string {
	orderedMapString := "OrderedMap["
	for e := om.list.Front(); e != nil; e = e.Next() {
		orderedMapString += fmt.Sprintf("%v", e.Value.(*listValue[K, V]))
		if e.Next() != nil {
			orderedMapString += " "
		}
	}
	return fmt.Sprintf("%s]", orderedMapString)
}

type listValue[K comparable, V any] struct {
	Key   K
	Value V
}

func (value *listValue[K, V]) String() string {
	return fmt.Sprintf("{%v: %v}", value.Key, value.Value)
}

func NewOrderedMap[K comparable, V any]() OrderedMap[K, V] {
	return &orderedMap[K, V]{
		m:    make(map[K]*list.Element),
		list: list.New(),
	}
}

func (om *orderedMap[K, V]) Set(key K, value V) {
	if elem, ok := om.m[key]; ok {
		elem.Value.(*listValue[K, V]).Value = value
		return
	}
	om.m[key] = om.list.PushBack(&listValue[K, V]{Key: key, Value: value})
}

func (om *orderedMap[K, V]) Get(key K) (V, bool) {
	if elem, ok := om.m[key]; ok {
		return elem.Value.(*listValue[K, V]).Value, true
	}
	var zero V
	return zero, false
}

func (om *orderedMap[K, V]) Remove(key K) {
	if elem, ok := om.m[key]; ok {
		om.list.Remove(elem)
		delete(om.m, key)
	}
}

func (om *orderedMap[K, V]) Keys() []K {
	keys := make([]K, 0, om.list.Len())
	for e := om.list.Front(); e != nil; e = e.Next() {
		keys = append(keys, e.Value.(*listValue[K, V]).Key)
	}
	return keys
}

func (om *orderedMap[K, V]) Values() []V {
	values := make([]V, 0, om.list.Len())
	for e := om.list.Front(); e != nil; e = e.Next() {
		values = append(values, e.Value.(*listValue[K, V]).Value)
	}
	return values
}

func (om *orderedMap[K, V]) Clear() {
	om.m = make(map[K]*list.Element)
	om.list.Init()
}

func (om *orderedMap[K, V]) Len() int {
	return om.list.Len()
}

func (om *orderedMap[K, V]) Iterator() <-chan listValue[K, V] {
	ch := make(chan listValue[K, V])
	go func() {
		for e := om.list.Front(); e != nil; e = e.Next() {
			ch <- listValue[K, V]{Key: e.Value.(*listValue[K, V]).Key, Value: e.Value.(*listValue[K, V]).Value}
		}
		close(ch)
	}()
	return ch
}
