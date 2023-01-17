package main

import (
	"fmt"
)

//请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
//
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
//key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
// 示例：
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	Store      map[int]*TwoListNode
	Head, Tail *TwoListNode
	Cap        int
}

type TwoListNode struct {
	Key  int
	Val  int
	Next *TwoListNode
	Pre  *TwoListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Store: make(map[int]*TwoListNode, capacity),
		Head:  nil,
		Tail:  nil,
		Cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Store[key]; ok {
		this.removeMid(node)
		this.addHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Store[key]; ok {
		this.removeMid(node)
		this.addHead(node)
		node.Val = value
		return
	}
	node := &TwoListNode{Key: key, Val: value}
	if len(this.Store) == this.Cap {
		this.removeTail()
	}
	this.addHead(node)
}

func (this *LRUCache) removeTail() {
	if this.Tail == nil {
		return
	}
	delete(this.Store, this.Tail.Key)
	if this.Head == this.Tail {
		this.Head = nil
		this.Tail = nil
		return
	}
	this.Tail.Pre.Next = nil
	temp := this.Tail.Pre
	this.Tail.Pre = nil
	this.Tail = temp
}

func (this *LRUCache) removeMid(node *TwoListNode) {
	if node == this.Head {
		delete(this.Store, node.Key)
		if this.Head == this.Tail {
			this.Head = nil
			this.Tail = nil
			return
		}
		if node.Next != nil {
			node.Next.Pre = nil
		}
		this.Head = node.Next
		node.Next = nil
	} else if node == this.Tail {
		this.removeTail()
	} else {
		delete(this.Store, node.Key)
		preNode := node.Pre
		nextNode := node.Next
		preNode.Next = nextNode
		nextNode.Pre = preNode
		node.Next = nil
		node.Pre = nil
	}
}

func (this *LRUCache) addHead(node *TwoListNode) {
	this.Store[node.Key] = node
	node.Next = this.Head
	if this.Head != nil {
		this.Head.Pre = node
	}
	this.Head = node
	if this.Tail == nil {
		this.Tail = node
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(1))
	cache.Put(5, 5)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(5))
}
