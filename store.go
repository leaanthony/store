package store

// Store is a struct that stores a generic value with getter and setters
type Store[T any] struct {
	subscriberCount int
	value           T
	subscribers     map[int]func(T)
}

// New creates a new store
func New[T any](initialValue T) *Store[T] {
	return &Store[T]{
		value:       initialValue,
		subscribers: make(map[int]func(T)),
	}
}

// Set sets the value of the store
func (s *Store[T]) Set(value T) {
	s.value = value
}

// Update updates the value of the store by running the given function
func (s *Store[T]) Update(updater func(T) T) {
	s.value = updater(s.value)
	for _, subscriber := range s.subscribers {
		subscriber(s.value)
	}
}

// Subscribe subscribes to the store
func (s *Store[T]) Subscribe(subscriber func(T)) func() {
	id := s.subscriberCount
	s.subscribers[id] = subscriber
	s.subscriberCount++
	return func() {
		delete(s.subscribers, id)
	}
}

// Get the value of the store
func (s *Store[T]) Get() T {
	return s.value
}
