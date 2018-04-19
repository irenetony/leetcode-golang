package main

import (
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"fmt"
)

func isSubtree(s *ds.TreeNode, t *ds.TreeNode) bool {
	if isSametree(s, t) {
		return true
	}

	if s == nil && t != nil || s != nil && t == nil {
		return false
	}

	if s != nil && t != nil {
		if s.Val == t.Val {
			if isSametree(s.Left, t.Left) && isSametree(s.Right, t.Right) {
				return true
			} else {
				if isSubtree(s.Left, t) {
					return true
				}
				if isSubtree(s.Right, t) {
					return true
				}
				return false
			}
		} else {
			if isSubtree(s.Left, t) {
				return true
			}
			if isSubtree(s.Right, t) {
				return true
			}
			return false
		}
	}

	return false
}

func isSametree(s *ds.TreeNode, t *ds.TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}
	if s != nil && t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	} else {
		return isSametree(s.Left, t.Left) && isSametree(s.Right, t.Right)
	}
}

func TestIsSubtree1() {
	s3 := &ds.TreeNode{Val: 3}
	s4 := &ds.TreeNode{Val: 4}
	s5 := &ds.TreeNode{Val: 5}
	s1 := &ds.TreeNode{Val: 1}
	s2 := &ds.TreeNode{Val: 2}

	s3.Left = s4
	s3.Right = s5
	s4.Left = s1
	s4.Right = s2

	t4 := &ds.TreeNode{Val: 4}
	t1 := &ds.TreeNode{Val: 1}
	t2 := &ds.TreeNode{Val: 2}

	t4.Left = t1
	t4.Right = t2

	fmt.Printf("%+v", isSubtree(s3, t4))
}

func TestIsSubtree2() {
	s3 := &ds.TreeNode{Val: 3}
	s4 := &ds.TreeNode{Val: 4}
	s5 := &ds.TreeNode{Val: 5}
	s1 := &ds.TreeNode{Val: 1}
	s2 := &ds.TreeNode{Val: 2}
	s0 := &ds.TreeNode{Val: 0}

	s3.Left = s4
	s3.Right = s5
	s4.Left = s1
	s4.Right = s2
	s2.Left = s0

	t4 := &ds.TreeNode{Val: 4}
	t1 := &ds.TreeNode{Val: 1}
	t2 := &ds.TreeNode{Val: 2}

	t4.Left = t1
	t4.Right = t2

	fmt.Printf("%+v\n", isSubtree(s3, t4))
}

func main() {
	TestIsSubtree1()
	TestIsSubtree2()
}
