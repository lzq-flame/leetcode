/**
2 * @Author: lzq
3 * @Date: 2021/9/9 20:43
4 */

//LRU缓存机制
//https://leetcode-cn.com/problems/lru-cache/
package main

type LRUCache struct {
	size       int
	capacity   int
	Cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{key: key, value: value}
}
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		Cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Cache[key]; !ok {
		return -1
	}
	node := this.Cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.Cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node := initDLinkedNode(key, value)
		this.Cache[key] = node
		this.size++
		this.addToHead(node)
		if this.size > this.capacity {
			tail := this.removeTail()
			delete(this.Cache, tail.key)
		}
	}

}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
