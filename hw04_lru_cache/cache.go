package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value any) bool
	Get(key Key) (any, bool)
	Clear()
}

type lruCache struct {
	mu       sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value any) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.items[key]
	if ok {
		c.items[key].Value = value
		c.queue.MoveToFront(c.items[key])
	} else {
		if c.queue.Len() == c.capacity {
			c.queue.Remove(c.queue.Back())
		}
		c.items[key] = c.queue.PushFront(value)
	}
	return ok
}

func (c *lruCache) Get(key Key) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, ok
	}
	c.queue.MoveToFront(item)
	return item.Value, true
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	// O(1)
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
