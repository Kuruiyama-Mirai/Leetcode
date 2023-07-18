package leetcode

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type TrieNode struct {
	nexts   [26]*TrieNode
	passCnt int
	isEnd   bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (t *Trie) Insert(word string) {
	if t.Search(word) {
		return
	}

	move := t.root
	for _, ch := range word {
		if move.nexts[ch-'a'] == nil {
			move.nexts[ch-'a'] = &TrieNode{}
		}
		move.nexts[ch-'a'].passCnt++
		move = move.nexts[ch-'a']
	}

	move.isEnd = true
}

func (t *Trie) Erase(word string) bool {
	if !t.Search(word) {
		return false
	}

	move := t.root
	for _, ch := range word {
		move.nexts[ch-'a'].passCnt--
		if move.nexts[ch-'a'].passCnt == 0 {
			move.nexts[ch-'a'] = nil
			return true
		}
		move = move.nexts[ch-'a']
	}

	move.isEnd = false
	return true
}

func (t *Trie) Search(word string) bool {
	node := t.search(word)
	//如果节点存在，并且节点的 isEnd 标识为 true，代表 word 存在
	return node != nil && node.isEnd
}

func (t *Trie) search(target string) *TrieNode {
	//移动指针从根节点出发
	move := t.root
	// 依次遍历 target 中的每个字符
	for _, ch := range target {
		// 倘若 nexts 中不存在对应于这个字符的节点，说明该单词没插入过，返回 nil
		if move.nexts[ch-'a'] == nil {
			return nil
		}
		move = move.nexts[ch-'a']
	}
	return move
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix) != nil
}

func (t *Trie) PassCnt(prefix string) int {
	node := t.search(prefix)
	if node == nil {
		return 0
	}
	return node.passCnt
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
