package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
	cap   int                   // 缓存容量
	cache map[int]*list.Element // 双向链表节点 指向的map
	list  *list.List            // 双向链表
}

type keyVal struct {
	key, val int // 节点的Key和Value
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*list.Element),
		list:  list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*keyVal).val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		elem.Value.(*keyVal).val = value
		return
	}
	if l.list.Len() >= l.cap {
		tail := l.list.Back()
		k := tail.Value.(*keyVal).key
		l.list.Remove(tail)
		delete(l.cache, k)
	}
	elem := l.list.PushFront(
		&keyVal{key, value},
	)
	l.cache[key] = elem
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
