package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        return &TreeNode{
            Val: val,
            Left: nil,
            Right: nil,
        }
    }
    layer := 1
    isDeep := false
    q := &Queue{}
    q.Push(root)
    for q.Len != 0 {
        count := q.Len
        if depth - 1 == layer {
            isDeep = true
        }
        for count != 0 {
            node := q.Pop()
            if isDeep {
                leftNode := node.Left
                newLeftNode := &TreeNode{Val: val}
                newLeftNode.Left = leftNode
                node.Left = newLeftNode
                rightNode := node.Right
                newRightNode := &TreeNode{Val: val}
                newRightNode.Right = rightNode
                node.Right = newRightNode
            } else {
                if node.Left != nil {
                    q.Push(node.Left)
                }
                if node.Right != nil {
                    q.Push(node.Right)
                }
            }
            count--
        }
        if isDeep {
            break
        } else {
            layer++
        }
    }
    return root
}

type Queue struct {
    Set []*TreeNode
    Len int
}

func (q *Queue) Push(node *TreeNode) {
    q.Set = append(q.Set, node)
    q.Len++
}

func (q *Queue) Pop() *TreeNode {
    if q.Len == 0 {
        return nil
    }
    node := q.Set[0]
    q.Set = q.Set[1:]
    q.Len--
    return node
}

func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    for i:=0; len(stack)>0; i++ {
        result = append(result, []int{})
        record := make([]*TreeNode, 0)
        for j:=0; j<len(stack); j++ {
            node := stack[j]
            result[i] = append(result[i], node.Val)
            if node.Left != nil {
                record = append(record, node.Left)
            }
            if node.Right != nil {
                record = append(record, node.Right)
            }
        }
        stack = record
    }
    return result
}

func main() {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
		},
	}
	root := addOneRow(p, 5, 4)
	fmt.Println(levelOrder(root))
}