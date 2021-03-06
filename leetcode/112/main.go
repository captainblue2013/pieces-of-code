package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
func main() {
	root := &TreeNode{
		Val: -2,
		Right: &TreeNode{
			Val: -3,
		},
	}
	fmt.Printf("%+v\n", []interface{}{hasPathSum(root, -5)})
}
