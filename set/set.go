package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] interface {
	Add(...T)
	Remove(...T)
	Clear()
	Contains(T) bool
	Len() int
	Slice() []T
	Equal(other Set[T]) bool
	Join(sep string) string
}

type set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(set[T])
	for _, item := range items {
		s.Add(item)
	}
	return s
}

func (s set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s set[T]) Clear() {
	for item := range s {
		delete(s, item)
	}
}

func (s set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s set[T]) Len() int {
	return len(s)
}

func (s set[T]) Slice() []T {
	slice := make([]T, 0, s.Len())
	for item := range s {
		slice = append(slice, item)
	}
	return slice
}

func (s set[T]) Equal(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for item := range s {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

func (s set[T]) Join(sep string) string {
	if s.Len() == 0 {
		return ""
	}

	var b strings.Builder
	first := true
	for item := range s {
		if !first {
			b.WriteString(sep)
		}
		fmt.Fprintf(&b, "%v", item)
		first = false
	}
	return b.String()
}

func (s set[T]) String() string {
	var b strings.Builder
	b.WriteString("Set{")
	b.WriteString(s.Join(", "))
	b.WriteString("}")
	return b.String()
}
