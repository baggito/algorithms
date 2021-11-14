package main

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var currentSum = 0
	return checkPath(root, currentSum, targetSum)
}

func checkPath(node *TreeNode, currentSum, targetSum int) bool {
	if node == nil {
		return false
	}

	currentSum += node.Val

	// if we meet a leaf
	if node.Left == nil && node.Right == nil {
		// check path sum so far
		if currentSum == targetSum {
			return true
		}
		return false
	}

	return checkPath(node.Right, currentSum, targetSum) ||
		checkPath(node.Left, currentSum, targetSum)
}
