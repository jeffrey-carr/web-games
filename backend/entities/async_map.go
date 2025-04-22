package entities

import (
	"maps"
	"sync"
)

// AsyncMap is a thread-safe map
type AsyncMap[K comparable, T any] struct {
	lock *sync.RWMutex
	data map[K]T
}

// NewAsyncMap creates a new async map
func NewAsyncMap[K comparable, T any](existingData map[K]T) AsyncMap[K, T] {
	// Copy the data into the local map
	d := map[K]T{}
	maps.Copy(d, existingData)

	return AsyncMap[K, T]{
		lock: &sync.RWMutex{},
		data: d,
	}
}

// Get gets the value associated with the specified key and
// a boolean representing if the value was found
func (m AsyncMap[K, T]) Get(key K) (T, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	item, ok := m.data[key]
	return item, ok
}

func (m AsyncMap[K, T]) Put(key K, item T) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = item
}

func (m AsyncMap[K, T]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.data, key)
}

func (m AsyncMap[K, T]) Size() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.data)
}
