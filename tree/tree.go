package tree

import (
	"fmt"
)

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("%d ,l=>%p,r=>%p", n.val, n.left, n.right)
}
func PreOrderTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println("tree node:", root.val)
	PreOrderTree(root.left)
	PreOrderTree(root.right)
}
func MiddleOrderTree(root *TreeNode) {
	if root == nil {
		return
	}
	MiddleOrderTree(root.left)
	fmt.Println("tree node:", root.val)
	MiddleOrderTree(root.right)
}
func PostOrderTree(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderTree(root.left)
	PostOrderTree(root.right)
	fmt.Println("tree node:", root.val)
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val: val}
	}
	if val < root.val {
		root.left = Insert(root.left, val)
		return root
	}
	root.right = Insert(root.right, val)
	return root
}

func Insert1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val: val}
	}
	if root.left == nil {
		root.left = Insert1(root.left, val)
	} else if root.right == nil {
		root.right = Insert1(root.right, val)
	} else {
		Insert1(root.left, val)
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
