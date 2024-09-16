package linkedlist

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	store    map[int]*Node
	left     *Node
	right    *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		store:    map[int]*Node{},
		left:     &Node{key: 0, val: 0},
		right:    &Node{key: 0, val: 0},
	}

	cache.left.next, cache.right.prev = cache.right, cache.left
	return cache
}

func (this *LRUCache) Remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

func (this *LRUCache) Insert(node *Node) {
	prev, next := this.right.prev, this.right
	next.prev = node
	prev.next = node
	node.next, node.prev = next, prev
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.store[key]; ok {
		this.Remove(this.store[key])
		this.Insert(this.store[key])
		return this.store[key].val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.store[key]; ok {
		this.Remove(this.store[key])
	}
	this.store[key] = &Node{key: key, val: value}
	this.Insert(this.store[key])

	if len(this.store) > this.capacity {
		// remove from the list and delete the LRU from hashmap
		lru := this.left.next
		this.Remove(lru)
		delete(this.store, lru.key)
	}
}
