package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：前序遍历+中序遍历，如果最后遍历的结果一样，则说明它们结构、数值都相同
// 方法二：深度遍历
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// pPreTraversal := preorderTraversal(p)
	// pInTraversal := inorderTraversal(p)
	// pLength := len(pPreTraversal)
	// qPreTraversal := preorderTraversal(q)
	// qInTraversal := inorderTraversal(q)
	// qLength := len(qPreTraversal)
	// if pLength != qLength {
	// 	return false
	// }
	// for i := 0; i < pLength; i++ {
	// 	if pPreTraversal[i] != qPreTraversal[i] {
	// 		return false
	// 	}
	// 	if pInTraversal[i] != qInTraversal[i] {
	// 		return false
	// 	}
	// }
	// return true

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func preorderTraversal(tree *TreeNode) []int {
	if tree == nil {
		return []int{0}
	}
	result := make([]int, 0)
	result = append(result, tree.Val)
	result = append(result, preorderTraversal(tree.Left)...)
	result = append(result, preorderTraversal(tree.Right)...)
	return result
}

func inorderTraversal(tree *TreeNode) []int {
	if tree == nil {
		return []int{0}
	}
	result := make([]int, 0)
	result = append(result, preorderTraversal(tree.Left)...)
	result = append(result, tree.Val)
	result = append(result, preorderTraversal(tree.Right)...)
	return result
}
