package leetcode

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := map[string]bool{}
	for _, w := range wordList {
		wordMap[w] = true
	}

	queue := []string{beginWord}
	level := 1
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return level
			}
			newWords := getCandidates(word)
			for _, newword := range newWords {
				if wordMap[newword] == true {
					queue = append(queue, newword)
					wordMap[newword] = false
				}
			}
		}
		level++
	}
	return 0
}

//替换26个字母，生成结果集
func getCandidates(word string) []string {
	res := make([]string, 0)

	for i := 0; i < len(word); i++ {
		for j := 'a'; j <= 'z'; j++ {
			res = append(res, word[:i]+string(j)+word[i+1:])
		}
	}
	return res
}

// @lc code=end
