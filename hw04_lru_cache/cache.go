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

func (c *lruCache) Set(key Key, value interface{}) bool {
	existsItem, ok := c.items[key]
	if ok {
		existsItem.Value = value
		c.queue.MoveToFront(existsItem)
		c.items[key] = existsItem

		return true
	}

	c.items[key] = c.queue.PushFront(value)

	if c.queue.Len() > c.capacity {
		back := c.queue.Back()
		var index Key
		for i, v := range c.items {
			if v == back {
				index = i
				break
			}
		}

		delete(c.items, index)
		c.queue.Remove(back)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	existsItem, ok := c.items[key]
	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(existsItem)

	return existsItem.Value, true
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
