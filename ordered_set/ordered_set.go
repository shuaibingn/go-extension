package ordered_set

import (
	"fmt"
	"strings"
)

type basevalue comparable

type OrderedSet[T basevalue] interface {
	Add(items ...T)
	Last() (T, int)
	Index(item T) int
	Remove(items ...T)
	Len() int
	Join(sep string) string
	Iterator() <-chan T
}

type orderedSet[T basevalue] struct {
	elems   []T
	elemMap map[T]int
}

func NewOrderedSet[T basevalue](items ...T) OrderedSet[T] {
	s := new(orderedSet[T])
	s.Add(items...)
	return s
}

func (s *orderedSet[T]) Add(items ...T) {
	if s.elemMap == nil {
		s.elemMap = make(map[T]int, len(items))
	}

	startIndex := len(s.elems)
	for _, item := range items {
		if _, ok := s.elemMap[item]; ok {
			continue
		}
		s.elems = append(s.elems, item)
		s.elemMap[item] = startIndex
		startIndex++
	}
}

func (s *orderedSet[T]) Last() (T, int) {
	if len(s.elems) == 0 {
		return *new(T), -1
	}
	lastElem := s.elems[len(s.elems)-1]
	return lastElem, s.elemMap[lastElem]
}

func (s *orderedSet[T]) Index(item T) int {
	if _, ok := s.elemMap[item]; !ok {
		return -1
	}
	return s.elemMap[item]
}

func (s *orderedSet[T]) Remove(items ...T) {
	toRemove := make(map[T]struct{})
	for _, item := range items {
		toRemove[item] = struct{}{}
	}

	newElems := make([]T, 0, len(s.elems))
	newElemMap := make(map[T]int)
	for _, elem := range s.elems {
		if _, ok := toRemove[elem]; !ok {
			newElems = append(newElems, elem)
			newElemMap[elem] = len(newElems) - 1
		}
	}

	s.elems = newElems
	s.elemMap = newElemMap
}

func (s *orderedSet[T]) Len() int {
	return len(s.elems)
}

func (s *orderedSet[T]) Join(sep string) string {
	if s.Len() == 0 {
		return ""
	}

	var b strings.Builder
	b.Grow(s.Len()*2 - 1)

	i := 0
	for _, item := range s.elems {
		b.WriteString(fmt.Sprintf("%v", item))
		i++
		if i < s.Len() {
			b.WriteString(sep)
		}
	}
	return b.String()
}

func (s *orderedSet[T]) Iterator() <-chan T {
	ch := make(chan T, len(s.elems))
	go func() {
		for _, item := range s.elems {
			ch <- item
		}
		close(ch)
	}()
	return ch
}

func (s *orderedSet[T]) String() string {
	var b strings.Builder
	b.WriteString("OrderedSet{")
	b.WriteString(s.Join(", "))
	b.WriteString("}")
	return b.String()
}
