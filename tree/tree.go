package tree

import (
	"fmt"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("%d ,l=>%p,r=>%p", n.Val, n.Left, n.Right)
}
func PreOrderTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println("tree node:", root.Val)
	PreOrderTree(root.Left)
	PreOrderTree(root.Right)
}
func MiddleOrderTree(root *TreeNode) {
	if root == nil {
		return
	}
	MiddleOrderTree(root.Left)
	fmt.Println("tree node:", root.Val)
	MiddleOrderTree(root.Right)
}
func PostOrderTree(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderTree(root.Left)
	PostOrderTree(root.Right)
	fmt.Println("tree node:", root.Val)
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
		return root
	}
	root.Right = Insert(root.Right, val)
	return root
}

func Insert1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Left == nil {
		root.Left = Insert1(root.Left, val)
	} else if root.Right == nil {
		root.Right = Insert1(root.Right, val)
	} else {
		Insert1(root.Left, val)
	}
	return root
}

func CreateBiTree() *TreeNode {
	var input int
	var root *TreeNode
	for {
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			break
		}
		root = Insert1(root, input)
	}
	return root
}
