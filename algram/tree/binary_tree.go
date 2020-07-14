// Package tree for
package tree

// TreeNode .
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	s := &Stack{}
	s.Push(root)
	for !s.Empty() {
		cur := s.Pop().(*TreeNode)
		res = append(res, cur.Val)
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
	return res
}

func dfsPreorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	s := &Stack{}
	s.Push(root)
	for !s.Empty() {
		cur := s.Pop().(*TreeNode)
		res = append(res, cur.Val)
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
	return res
}

func dfsPostorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	s := &Stack{}
	s.Push(root)
	for !s.Empty() {
		cur := s.Pop().(*TreeNode)
		res = append(res, cur.Val)
		if cur.Left != nil {
			s.Push(cur.Left)
		}
		if cur.Right != nil {
			s.Push(cur.Right)
		}
	}
	for i, j := 0, len(res)-1; i <= j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	s := &Stack{}
	cur := root
	for !s.Empty() || cur != nil {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		cur = s.Pop().(*TreeNode)
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	s := &Stack{}
	cur := root
	for !s.Empty() || cur != nil {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		cur = s.Pop().(*TreeNode)
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

type Stack struct {
	Elem []interface{}
}

func (s *Stack) Pop() interface{} {
	length := len(s.Elem)
	if length == 0 {
		return nil
	}
	e := s.Elem[length-1]
	s.Elem = s.Elem[:length-1]
	return e
}

func (s *Stack) Peek() interface{} {
	length := len(s.Elem)
	if length == 0 {
		return nil
	}
	return s.Elem[length-1]
}

func (s *Stack) Empty() bool {
	if len(s.Elem) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(e interface{}) {
	s.Elem = append(s.Elem, e)
}
