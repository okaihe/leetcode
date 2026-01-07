package main

import "fmt"

func main() {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.PrintNodeList()
	obj.Put(1, 1)
	obj.PrintNodeList()
	obj.Put(2, 3)
	obj.PrintNodeList()
	obj.Put(4, 1)
	obj.PrintNodeList()
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}

func (this *LRUCache) PrintNodeList() {
	if this.startNode != nil {
		node := this.startNode
		fmt.Print(node)
		for node.next != nil {
			node = node.next
			fmt.Print("", node)
		}
	}
	fmt.Println()
}

type LRUCache struct {
	capacity        int
	currentCapacity int
	startNode       *Node
	endNode         *Node
	hashMap         map[int]*Node
}

type Node struct {
	next  *Node
	last  *Node
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:        capacity,
		currentCapacity: 0,
		startNode:       nil,
		endNode:         nil,
		hashMap:         make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	node, doesExist := this.hashMap[key]
	if doesExist {
		// node is first node
		if node.last == nil {
			return node.value
		}

		// node is last node
		if node.next == nil {
			node.last.next = nil
			this.endNode = node.last
		} else {
			// node is middle node
			node.last.next = node.next
			node.next.last = node.last
		}

		node.last = nil
		node.next = this.startNode
		this.startNode.last = node
		this.startNode = node
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, doesExist := this.hashMap[key]
	if doesExist {
		// node is first node
		if node.last == nil {
			node.value = value
			return
		}

		// node is last node
		if node.next == nil {
			node.last.next = nil
			this.endNode = node.last
		} else {
			// node is middle node
			node.last.next = node.next
			node.next.last = node.last
		}

		node.last = nil
		node.next = this.startNode
		this.startNode.last = node
		this.startNode = node
		this.startNode.value = value
	} else {
		var newNode *Node = &Node{key: key, value: value}

		if this.currentCapacity != 0 {
			this.startNode.last = newNode
		}

		newNode.next = this.startNode
		this.startNode = newNode

		if this.currentCapacity == 0 {
			this.endNode = newNode
		}

		this.hashMap[key] = newNode

		if this.currentCapacity == this.capacity {
			lastNode := this.endNode
			this.endNode.last.next = nil
			this.endNode = this.endNode.last
			delete(this.hashMap, lastNode.key)
		} else {
			this.currentCapacity += 1
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
