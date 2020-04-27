package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/0xAX/go-algorithms/stack"
)

func compressString(S string) string {
	var lastChar int32
	var cnt int
	var out string
	for _, v := range S {
		if lastChar == 0 {
			out += string(v)
			lastChar = v
			cnt = 1

		} else if lastChar == v {
			cnt++
		} else {
			lastChar = v
			out += strconv.FormatInt(int64(cnt), 10) + string(v)
			cnt = 1
		}
	}
	out += strconv.FormatInt(int64(cnt), 10)
	if len(out) > len(S) {
		return S
	}
	return out
}

func CountCharactersGood(words []string, chars string) int {
	var length int
outloop:
	for _, v := range words {
		for _, v1 := range v {
			if strings.Count(v, string(v1)) > strings.Count(chars, string(v1)) {
				continue outloop
			}
		}
		length += len(v)

	}
	return length
}

func CountCharactersBad(words []string, chars string) int {
	cnt := make(map[rune]int, 26)
	var length int
	for _, v := range chars {
		cnt[v]++
	}
outloop:
	for _, v := range words {
		wordCnt := make(map[rune]int)
		for _, iv := range v {
			wordCnt[iv]++
		}
		for k, iiv := range wordCnt {
			if iiv > cnt[k] {
				continue outloop
			}
		}
		length += len(v)
	}
	return length
}

func Loop() map[int]int {
	a := make(map[int]int, 100)
	return a
}

// 836
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return len(rec1) == 4 && len(rec2) == 4 &&
		(!(rec1[2] < rec2[0] ||
			rec1[0] > rec2[2] ||
			rec1[3] < rec2[1] ||
			rec1[1] > rec2[3]))
}

// 409
func longestPalindrome(s string) int {
	var pair int
	cnt := make([]int, 26*2)
	for _, v := range s {
		cnt[v-'A']++
	}
	for _, v := range cnt {
		pair += 2 * (v / 2)
		if v%2 == 1 && pair%2 == 0 {
			pair++
		}
	}
	return pair
}

// 133
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

var hashMap = make(map[int]*Node, 110)

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := hashMap[node.Val]; ok {
		return v
	}
	newPtr := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 1),
	}
	hashMap[newPtr.Val] = newPtr
	for _, v := range node.Neighbors {
		newPtr.Neighbors = append(newPtr.Neighbors, cloneGraph(v))
	}
	return newPtr
}

/////////// 876
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var cnt int
	ptr := head
	for {
		if ptr == nil {
			break
		}
		cnt++
		ptr = ptr.Next
	}
	aim := cnt/2 + 1
	var aimPtr = head
	for i := 0; i < aim-1; i++ {
		aimPtr = aimPtr.Next
	}
	return aimPtr
}

/////10

func isMatch(s string, p string) bool {

	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := (len(s) != 0) && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		if firstMatch {
			return isMatch(substr(s, 1), p)
		}
		return isMatch(s, p[2:])

	}
	return isMatch(substr(s, 1), p[1:])
}
func substr(s string, num int) string {
	if len(s) <= num {
		return ""
	}
	return s[num:]
}

//面试题 17.16. 按摩师
// 递推方程
// dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
func massage(nums []int) int {
	l := len(nums)
	last2 := 0
	last1 := 0
	var now int
	for i := 0; i < l; i++ {
		now = max(last1, last2+nums[i])
		last2 = last1
		last1 = now
	}
	return now
}

// 860 柠檬水
func lemonadeChange(bills []int) bool {
	l := make([]int, 4)
	for _, v := range bills {
		if v == 5 {
			l[0]++
			continue
		}
		if l[0]*5+l[1]*10 >= v-5 && l[0] != 0 {
			l[v/5-1]++
			if v == 20 {
				if l[1] >= 1 {
					l[1]--
					l[0]--
				} else {
					if l[0] >= 3 {
						l[0] -= 3
					} else {
						return false
					}
				}

			} else if v == 10 {
				l[0]--
			}
			continue
		}
		return false
	}
	return true
}

// 892. 三维形体的表面积
func surfaceArea(grid [][]int) int {
	var total int
	length := len(grid)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if grid[i][j] > 0 {
				total += 2 + grid[i][j]*4
			} else {
				continue
			}
			if i > 1 {
				if grid[i][j] > grid[i-1][j] {
					total -= grid[i-1][j]
				} else {
					total -= grid[i][j]
				}
			}
			if j > 1 {
				if grid[i][j] > grid[i][j-1] {
					total -= grid[i][j-1]
				} else {
					total -= grid[i][j]
				}
			}
			if i < length-1 {
				if grid[i][j] > grid[i+1][j] {
					total -= grid[i+1][j]
				} else {
					total -= grid[i][j]
				}
			}
			if j < length-1 {
				if grid[i][j] > grid[i][j+1] {
					total -= grid[i][j+1]
				} else {
					total -= grid[i][j]
				}
			}
		}
	}
	return total
}

// 面试题 02.05. 链表求和
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := new(ListNode)
	tmp1 := l1
	tmp2 := l2
	tmp3 := sum
	var a int
	for {
		if tmp1 == nil && tmp2 == nil && a == 0 {
			return sum
		}
		tmp3.Val += a
		if tmp1 != nil {
			tmp3.Val += tmp1.Val
			tmp1 = tmp1.Next
		}
		if tmp2 != nil {
			tmp3.Val += tmp2.Val
			tmp2 = tmp2.Next
		}
		a = tmp3.Val / 10
		tmp3.Val = tmp3.Val % 10
		if tmp1 != nil || tmp2 != nil || a != 0 {
			tmp3.Next = new(ListNode)
		}
		tmp3 = tmp3.Next
	}
	return tmp3
}

// 914
func hasGroupsSizeX(deck []int) bool {
	var cnt = make(map[int]int)
	for _, v := range deck {
		cnt[v]++
	}
BC:
	for i := 2; i <= len(deck); i++ {
		for _, v := range cnt {
			if v/i >= 1 && v%i == 0 {
				fmt.Printf("count:%d,yueshu:%d", v, i)
				continue
			}
			break BC
		}
		return true
	}
	return false
}

// 1035 最长公共子序列 LCS
// f(x,y)表示 A x长度，b y长度最长子序列
// f(x,y) = max{f(x-1,y-1)+sum(x,y),f(x-1,y),f(x,y-1)}
// f(1,1) = sum(1,1)
func maxUncrossedLines(A []int, B []int) int {
	lenA := len(A)
	lenB := len(B)
	c := make([][]int, lenA+1)
	for i := 0; i < lenA+1; i++ {
		c[i] = make([]int, lenB+1)
	}
	for i := 0; i < lenA; i++ {
		for j := 0; i < lenB; j++ {
			if A[i] == B[j] {
				c[i+1][j+1] = c[i][j] + 1
			} else {
				c[i+1][j+1] = max(c[i][j+1], c[i+1][j])
			}
		}
	}
	return c[lenA][lenB]
}

func sum(x, y int) int {
	if x == y {
		return 1
	}
	return 0
}

// 72
// 给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

// 你可以对一个单词进行如下三种操作：

// 插入一个字符
// 删除一个字符
// 替换一个字符

func minDistance(word1 string, word2 string) int {
	var rec = make([][]int, len(word1))
	len2 := len(word2)
	for i := range word1 {
		rec[i] = make([]int, len2)
		for j := range word2 {
			rec[i][j] = -1
		}
	}
	return dp(word1, word2, rec)
}

func dp(word1 string, word2 string, rec [][]int) int {
	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 || len2 == 0 {
		return max(len1, len2)
	}
	if rec[len1-1][len2-1] != -1 {
		return rec[len1-1][len2-1]
	}
	var cost int
	if word1[len1-1] == word2[len2-1] {
		cost = 0
	} else {
		cost = 1
	}
	a := dp(word1[:len1-1], word2[:len2], rec) + 1
	b := dp(word1[:len1], word2[:len2-1], rec) + 1
	c := dp(word1[:len1-1], word2[:len2-1], rec) + cost
	res := min(a, b, c)
	rec[len1-1][len2-1] = res
	return res
}

// 289
func gameOfLife(board [][]int) {
	lenX := len(board)
	lenY := len(board[0])
	var dir = [8][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	for k1, v1 := range board {
		for k2, v2 := range v1 {
			sum := 0
			for _, vdir := range dir {
				x := k1 + vdir[0]
				y := k2 + vdir[1]
				if x >= 0 && x < lenX && y >= 0 && y < lenY {
					sum += board[x][y] & 1
				}
			}
			v := v2 & 1
			if sum == 3 {
				board[k1][k2] |= 1 << 1
			} else if sum == 2 {
				board[k1][k2] |= v << 1
			} else {
				board[k1][k2] |= 0 << 1
			}
		}
	}
	for k1, v1 := range board {
		for k2 := range v1 {
			board[k1][k2] >>= 1
		}
	}
}

// common func
func max(a ...int) int {
	var max int
	for _, v := range a {
		if max < v {
			max = v
		}
	}
	return max
}

func min(a ...int) int {
	min := a[0]
	for _, v := range a {
		if min > v {
			min = v
		}
	}
	return min
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	st := stack.New()
	res := []int{}
	for {
		if v := st.Pop(); v != nil {
			res = append(res, v.(*TreeNode).Val)
			st.Push(v.(*TreeNode).Left)
			st.Push(v.(*TreeNode).Right)
		} else {
			break
		}
	}
	return res
}

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	dfs("", 2*n, n, n)
	return res
}

func dfs(current string, max, left, right int) {
	if check(current) {
		fmt.Println(current)
		if len(current) == max {
			fmt.Println(current, max)
			res = append(res, current)
			return
		}
		if left > 0 {
			dfs(current+"(", max, left-1, right)
		}
		if right > 0 {
			dfs(current+")", max, left, right-1)
		}
	}
}

func check(s string) bool {
	if strings.Count(s, "(") >= strings.Count(s, ")") {
		return true
	}
	return false
}

var dMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}
	var ans []string
	if len(digits) == 1 {
		for _, v := range dMap[digits[0]] {
			ans = append(ans, string(v))
		}
	}
	lastString := letterCombinations(digits[:len(digits)-1])
	for _, v := range dMap[digits[len(digits)-1]] {
		for _, v1 := range lastString {
			ans = append(ans, v1+string(v))
		}
	}
	return ans
}

type MyHashMap struct {
	h [][]*KV
}

type KV struct {
	K int
	V int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{h: make([][]*KV, 1000)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	mod := key % 1000
	for _, v := range this.h[mod] {
		if v.K == key {
			v.V = value
			return
		}
	}
	this.h[mod] = append(this.h[mod], &KV{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	var res int = -1
	for _, v := range this.h[key%1000] {
		if v != nil && v.K == key {
			res = v.V
			break
		}
	}
	return res
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	for k, v := range this.h[key%1000] {
		if v != nil && v.K == key {
			this.h[key%1000][k] = nil
			break
		}
	}
}

func waysToChange(n int) int {
	if n == 0 {
		return 1
	}
	if n >= 25 {
		return waysToChange(n-25) + waysToChange(n-10) + waysToChange(n-5) + waysToChange(n-1)
	}
	if n >= 10 {
		return waysToChange(n-10) + waysToChange(n-5) + waysToChange(n-1)
	}

	if n >= 5 {
		return waysToChange(n-5) + waysToChange(n-1)
	}
	return waysToChange(n - 1)
}
