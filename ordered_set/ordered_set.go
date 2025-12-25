package ordered_set

import (
	"fmt"
	"strings"
)

type baseValue comparable

type OrderedSet[T baseValue] interface {
	Add(items ...T)
	Remove(items ...T)
	Clear()
	Contains(item T) bool
	Index(item T) int
	Get(index int) (T, bool)
	First() (T, bool)
	Last() (T, bool)
	Len() int
	Join(sep string) string
	Slice() []T
	SliceRef() []T
	ForEach(fn func(index int, item T) bool)
}

type orderedSet[T baseValue] struct {
	elems   []T
	elemMap map[T]struct{}
}

func NewOrderedSet[T baseValue](items ...T) OrderedSet[T] {
	s := new(orderedSet[T])
	s.Add(items...)
	return s
}

func (s *orderedSet[T]) Add(items ...T) {
	if s.elemMap == nil {
		s.elemMap = make(map[T]struct{}, len(items))
	}

	for _, item := range items {
		if _, ok := s.elemMap[item]; ok {
			continue
		}
		s.elems = append(s.elems, item)
		s.elemMap[item] = struct{}{}
	}
}

func (s *orderedSet[T]) Clear() {
	s.elems = nil
	s.elemMap = nil
}

func (s *orderedSet[T]) Contains(item T) bool {
	_, ok := s.elemMap[item]
	return ok
}

func (s *orderedSet[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(s.elems) {
		return *new(T), false
	}
	return s.elems[index], true
}

func (s *orderedSet[T]) First() (T, bool) {
	if len(s.elems) == 0 {
		return *new(T), false
	}
	return s.elems[0], true
}

func (s *orderedSet[T]) Last() (T, bool) {
	if len(s.elems) == 0 {
		return *new(T), false
	}
	return s.elems[len(s.elems)-1], true
}

func (s *orderedSet[T]) Index(item T) int {
	if _, ok := s.elemMap[item]; !ok {
		return -1
	}

	for i, elem := range s.elems {
		if elem == item {
			return i
		}
	}
	return -1
}

func (s *orderedSet[T]) Remove(items ...T) {
	if len(items) == 0 {
		return
	}

	for _, item := range items {
		delete(s.elemMap, item)
	}

	n := 0
	for _, elem := range s.elems {
		if _, ok := s.elemMap[elem]; ok {
			s.elems[n] = elem
			n++
		}
	}

	// 清除残留引用，防止内存泄漏
	var zero T
	for i := n; i < len(s.elems); i++ {
		s.elems[i] = zero
	}
	s.elems = s.elems[:n]
}

func (s *orderedSet[T]) Len() int {
	return len(s.elems)
}

func (s *orderedSet[T]) Join(sep string) string {
	if s.Len() == 0 {
		return ""
	}

	var b strings.Builder
	for i, item := range s.elems {
		if i > 0 {
			b.WriteString(sep)
		}
		_, _ = fmt.Fprintf(&b, "%v", item)
	}
	return b.String()
}

func (s *orderedSet[T]) ForEach(fn func(index int, item T) bool) {
	for i, item := range s.elems {
		if !fn(i, item) {
			break
		}
	}
}

func (s *orderedSet[T]) Slice() []T {
	result := make([]T, len(s.elems))
	copy(result, s.elems)
	return result
}

// SliceRef returns the internal slice reference(zero copy, O(1))
// Warning: Don't modify the returned slice, it will break the Set's data consistency
// If you need to modify, use Slice() to get a copy
func (s *orderedSet[T]) SliceRef() []T {
	return s.elems
}

func (s *orderedSet[T]) String() string {
	var b strings.Builder
	b.WriteString("OrderedSet{")
	b.WriteString(s.Join(", "))
	b.WriteString("}")
	return b.String()
}
