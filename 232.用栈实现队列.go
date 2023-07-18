package leetcode

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	s1, s2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.s1 = append(q.s1, x)
}

/**
 * 删除队头的元素并返回
 */
func (q *MyQueue) Pop() int {
	q.Peek()
	val := q.s2[len(q.s2)-1]
	q.s2 = q.s2[:len(q.s2)-1]
	return val
}

func (q *MyQueue) Peek() int {
	if len(q.s2) == 0 {
		// 把 s1 元素压入 s2
		for len(q.s1) != 0 {
			q.s2 = append(q.s2, q.s1[len(q.s1)-1])
			q.s1 = q.s1[:len(q.s1)-1]
		}
	}
	return q.s2[len(q.s2)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.s1) == 0 && len(q.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
