package main

import (
	"container/list"
)

type CacheNode struct {
	key   int
	value int
	freq  int
}

type LFUCache struct {
	capacity        int
	currentCapacity int
	minFrequency    int
	frequencyMap    map[int]*list.List
	nodeMap         map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:        capacity,
		currentCapacity: 0,
		minFrequency:    0,
		frequencyMap:    make(map[int]*list.List),
		nodeMap:         make(map[int]*list.Element),
	}
}

func (this *LFUCache) Get(key int) int {
	elem, ok := this.nodeMap[key]
	if !ok {
		return -1
	}

	node := elem.Value.(*CacheNode)

	oldFreq := node.freq

	this.frequencyMap[oldFreq].Remove(elem)
	if oldFreq == this.minFrequency && this.frequencyMap[oldFreq].Len() == 0 {
		this.minFrequency++
	}

	node.freq++

	listForFreq, ok := this.frequencyMap[node.freq]
	if !ok {
		listForFreq = list.New()
		this.frequencyMap[node.freq] = listForFreq
	}

	newElem := listForFreq.PushFront(node)
	this.nodeMap[key] = newElem

	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if elem, ok := this.nodeMap[key]; ok {
		node := elem.Value.(*CacheNode)
		oldFreq := node.freq

		this.frequencyMap[oldFreq].Remove(elem)
		if oldFreq == this.minFrequency && this.frequencyMap[oldFreq].Len() == 0 {
			this.minFrequency++
		}

		node.value = value
		node.freq++

		listForFreq, ok := this.frequencyMap[node.freq]
		if !ok {
			listForFreq = list.New()
			this.frequencyMap[node.freq] = listForFreq
		}

		newElem := listForFreq.PushFront(node)
		this.nodeMap[key] = newElem
		return
	}

	if this.currentCapacity == this.capacity {
		minFreqList := this.frequencyMap[this.minFrequency]
		lfuNode := minFreqList.Back()
		lfuCacheNode := lfuNode.Value.(*CacheNode)
		minFreqList.Remove(lfuNode)
		delete(this.nodeMap, lfuCacheNode.key)
	} else {
		this.currentCapacity += 1
	}

	node := &CacheNode{
		key:   key,
		value: value,
		freq:  1,
	}

	listForFreq, ok := this.frequencyMap[1]
	if !ok {
		listForFreq = list.New()
		this.frequencyMap[1] = listForFreq
	}

	elem := listForFreq.PushFront(node)
	this.nodeMap[key] = elem
	this.minFrequency = 1
}
