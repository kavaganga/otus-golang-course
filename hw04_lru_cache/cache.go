package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value any) bool
	Get(key Key) (any, bool)
	Clear()
}

// Version sync.Map
/*
type lruCache struct {
	capacity int
	queue    List
	items    sync.Map
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    sync.Map{},
	}
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items.Range(func(key, value any) bool {
		c.items.Delete(key)
		return true
	})
}

func (c *lruCache) Get(key Key) (any, bool) {
	item, ok := c.items.Load(key)
	if !ok {
		return nil, ok
	}
	c.queue.MoveToFront(item.(*ListItem))
	return item.(*ListItem).Value, true
}

func (c *lruCache) Set(key Key, value any) bool {
	v := value.(*ListItem)
	_, ok := c.items.Swap(key, v)

	if ok {
		c.queue.MoveToFront(v)
	} else {
		c.queue.PushFront(v)
	}
	if c.queue.Len() > c.capacity {
		c.queue.Remove(c.queue.Back())
	}
	return ok
}
*/
// Version map + mutex lock

type lruCache struct {
	mu       sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
	keys     map[*ListItem]Key
	//items sync.Map
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		keys:     make(map[*ListItem]Key, capacity),
		//items: sync.Map{},
	}
}

func (c *lruCache) Set(key Key, value any) bool {
	c.mu.Lock()
	//log.Println("SET", key)
	defer c.mu.Unlock()

	_, ok := c.items[key]
	if ok {
		c.items[key].Value = value
		c.queue.MoveToFront(c.items[key])
	} else {
		if c.queue.Len() == c.capacity {
			val := c.queue.Back()
			c.queue.Remove(val)
			old_key := c.keys[val]
			delete(c.items, old_key)
			delete(c.keys, val)
		}
		c.items[key] = c.queue.PushFront(value)
	}
	c.keys[c.queue.Front()] = key
	return ok
}

func (c *lruCache) Get(key Key) (any, bool) {
	c.mu.RLock()
	//log.Println("GET", key)
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
	c.keys = make(map[*ListItem]Key, c.capacity)
}
