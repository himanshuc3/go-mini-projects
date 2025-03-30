package main

import (
	"slices"
	"sync"
	"time"
)

// 1. What is the diff b/w Time and Duration types?
type entryWithTimeout[V any] struct {
	value   V
	expires time.Time
}

// NOTE:
// 1. RWMutex is generally faster
// than Mutex but mutex is used more often
type Cache[K comparable, V any] struct {
	data              map[K]entryWithTimeout[V]
	mu                sync.Mutex
	ttl               time.Duration
	maxSize           int
	chronologicalKeys []K
}

func New[K comparable, V any](maxSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		// NOTE:
		// 1. Note giving it any memory/size upfront
		data:              make(map[K]entryWithTimeout[V]),
		ttl:               ttl,
		maxSize:           maxSize,
		chronologicalKeys: make([]K, 0, maxSize),
	}
}

// NOTE:
// 1. Testing this function initially can be mapped
// to 1:1 to the behavior of testing map structure
// itself, which is an implementation already tested
// by go divas
// 2. Fetching packages via: go get -v package_name
func (c *Cache[K, V]) Read(key K) (V, bool) {
	// NOTE:
	// 1. Read lock is not Mutually exclusive with another
	// read lock, but exclusive with a write/normal lock.
	c.mu.Lock()
	defer c.mu.Unlock()

	var zeroV V

	e, ok := c.data[key]

	switch {
	case !ok:
		// NOTE:
		// 1. Why can't we return V{}
		return zeroV, false
	case e.expires.Before(time.Now()):
		c.deleteKeyValue(key)
		return zeroV, false
	default:
		return e.value, true
	}
}

func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, alreadyPresent := c.data[key]

	switch {
	case alreadyPresent:
		c.deleteKeyValue(key)
	case len(c.data) == c.maxSize:
		c.deleteKeyValue(c.chronologicalKeys[0])
	}

	c.addKeyValue(key, value)
	return nil
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.deleteKeyValue(key)
}

func (c *Cache[K, V]) addKeyValue(key K, value V) {
	c.data[key] = entryWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	c.chronologicalKeys = append(c.chronologicalKeys, key)
}

func (c *Cache[K, V]) deleteKeyValue(key K) {
	// NOTE:
	// 1. Much Much efficient O(n) deletions
	c.chronologicalKeys = slices.DeleteFunc(c.chronologicalKeys, func(k K) bool {
		return k == key
	})
	delete(c.data, key)
}
