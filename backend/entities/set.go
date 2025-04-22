package entities

import (
	"web_games/utils"
)

// Set implements a standard set
type Set[T comparable] interface {
	Add(item T)
	Has(item T) bool
	Delete(item T)
	Slice() []T
	Size() int
}

// NewSet creates a new set with the items provided
func NewSet[T comparable](baseItems []T) Set[T] {
	newSet := set[T]{
		items: map[T]struct{}{},
	}

	for _, item := range baseItems {
		newSet.Add(item)
	}

	return &newSet
}

type set[T comparable] struct {
	items map[T]struct{}
}

func (s *set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

func (s *set[T]) Has(item T) bool {
	_, exists := s.items[item]
	return exists
}

func (s *set[T]) Delete(item T) {
	delete(s.items, item)
}

func (s *set[T]) Slice() []T {
	allItems := []T{}
	for key := range s.items {
		allItems = append(allItems, key)
	}

	return allItems
}

func (s *set[T]) Size() int {
	return len(s.items)
}

// SetStack is a stack that can only contain one instance of each item
type SetStack[T comparable] interface {
	Push(item T)
	Pop() (T, bool)
	Peek() (T, bool)
	Has(item T) bool
	Size() int
	Slice() []T
}

type setStack[T comparable] struct {
	// Hmm, do we care about space or time?
	items Set[T]
	order []T
}

// NewSetStack returns a new SetStack
func NewSetStack[T comparable](baseItems ...T) SetStack[T] {
	items := NewSet(baseItems)
	var order []T
	copy(order, baseItems)

	return &setStack[T]{items: items, order: order}
}

func (s *setStack[T]) Push(item T) {
	if !s.items.Has(item) {
		s.items.Add(item)
		s.order = append(s.order, item)
	}
}

func (s *setStack[T]) Pop() (T, bool) {
	if s.items.Size() == 0 {
		return utils.Zero[T](), false
	}
	if len(s.order) == 0 {
		return utils.Zero[T](), false
	}

	lastItem := s.order[len(s.order)-1]
	s.order = s.order[:len(s.order)-1]
	s.items.Delete(lastItem)

	return lastItem, true
}

func (s *setStack[T]) Peek() (T, bool) {
	if s.items.Size() == 0 {
		return utils.Zero[T](), false
	}

	return s.order[len(s.order)-1], true
}

func (s *setStack[T]) Has(item T) bool {
	return s.items.Has(item)
}

func (s *setStack[T]) Size() int {
	return len(s.order)
}

func (s *setStack[T]) Slice() []T {
	return s.order
}
