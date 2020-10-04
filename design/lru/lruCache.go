package lru

type Cache struct {
	size       int
	cap        int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, val   int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, val int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		val: val,
	}
}

func NewLRUCache(cap int) Cache {
	c := Cache{
		cache: map[int]*DLinkedNode{},
		head:  initDLinkedNode(0, 0), // dummy node
		tail:  initDLinkedNode(0, 0), // dummy node
		cap:   cap,
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (c *Cache) Get(key int) int {
	if _, ok := c.cache[key]; !ok {
		return -1
	}
	node := c.cache[key]
	// move this node to the head of the list
	c.moveToHead(node)
	return node.val
}

func (c *Cache) Put(key, val int) {
	if _, ok := c.cache[key]; !ok {
		// add a new node to the cache map
		node := initDLinkedNode(key, val)
		if c.size == c.cap {
			toDel := c.removeTail()
			delete(c.cache, toDel.key)
			c.size--
		}
		c.size++
		c.cache[key] = node
		c.addHead(node)
	} else {
		// the key has already existed
		node := c.cache[key]
		node.val = val
		c.moveToHead(node)
	}
}

func (c *Cache) addHead(node *DLinkedNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *Cache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *Cache) moveToHead(node *DLinkedNode) {
	c.removeNode(node)
	c.addHead(node)
}

func (c *Cache) removeTail() *DLinkedNode {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
