package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLayer := 1
	if root.Left == nil && root.Right == nil {
		return maxLayer
	}
	maxResult := root.Val
	travel := make([]*TreeNode, 0)
	travel = append(travel, root)
	for i := 1; len(travel) > 0; i++ {
		sum := 0
		length := len(travel)
		for j := 0; j < length; j++ {
			sum += travel[j].Val
			node := travel[j]
			if node.Left != nil {
				travel = append(travel, node.Left)
			}
			if node.Right != nil {
				travel = append(travel, node.Right)
			}
		}
		travel = travel[length:]
		if sum > maxResult {
			maxResult = sum
			maxLayer = i
		}
	}
	return maxLayer
}
