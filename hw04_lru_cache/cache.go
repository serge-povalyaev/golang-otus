package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if c.capacity == 0 {
		return false
	}

	existsItem, ok := c.items[key]
	if ok {
		cacheItem := existsItem.Value.(cacheItem)
		cacheItem.value = value
		existsItem.Value = cacheItem
		c.queue.MoveToFront(existsItem)
		c.items[key] = existsItem

		return true
	}

	if c.queue.Len() == c.capacity {
		back := c.queue.Back()
		delete(c.items, back.Value.(cacheItem).key)
		c.queue.Remove(back)
	}

	c.items[key] = c.queue.PushFront(cacheItem{key, value})

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	existsItem, ok := c.items[key]
	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(existsItem)

	return existsItem.Value.(cacheItem).value, true
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
