package sync

import (
	"sync"
)

// Store is a struct that stores a generic value with getter and setters
type Store[T any] struct {
	subscriberCount int
	value           T
	subscribers     map[int]func(T)
	mutex           sync.RWMutex
}

// New creates a new thread-safe store
func New[T any](initialValue T) *Store[T] {
	return &Store[T]{
		value:       initialValue,
		subscribers: make(map[int]func(T)),
	}
}

// Set sets the value of the store
func (s *Store[T]) Set(value T) {
	s.mutex.Lock()
	s.value = value
	s.mutex.Unlock()
}

// Update updates the value of the store by running the given function
func (s *Store[T]) Update(updater func(T) T) {
	s.mutex.Lock()
	s.value = updater(s.value)
	for _, subscriber := range s.subscribers {
		subscriber(s.value)
	}
	s.mutex.Unlock()
}

// Subscribe subscribes to the store
func (s *Store[T]) Subscribe(subscriber func(T)) func() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	id := s.subscriberCount
	s.subscribers[id] = subscriber
	s.subscriberCount++
	return func() {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		delete(s.subscribers, id)
	}
}

// Get the value of the store
func (s *Store[T]) Get() T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.value
}
