package main

import (
	"math"
)

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, int64(math.MinInt64), int64(math.MaxInt64))
}

func validate(root *TreeNode, low int64, high int64) bool {
	if root == nil {
		return true
	}

	if (int64(root.Val) <= low) || (int64(root.Val) >= high) {
		return false
	}

	return validate(root.Left, low, int64(root.Val)) &&  validate(root.Right, int64(root.Val), high)
}