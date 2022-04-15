package tree

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Node tree节点
type Node[T constraints.Ordered] struct {
	Val         T
	Left, Right *Node[T]
}

func (n *Node) String() string {
	return fmt.Sprintf("%v ,l=>%p,r=>%p", n.Val, n.Left, n.Right)
}
func PreOrderTree[T constraints.Ordered](root *Node[T]) {
	if root == nil {
		return
	}
	fmt.Println("tree node:", root.Val)
	PreOrderTree(root.Left)
	PreOrderTree(root.Right)
}
func MiddleOrderTree[T constraints.Ordered](root *Node[T]) {
	if root == nil {
		return
	}
	MiddleOrderTree(root.Left)
	fmt.Println("tree node:", root.Val)
	MiddleOrderTree(root.Right)
}
func PostOrderTree[T constraints.Ordered](root *Node[T]) {
	if root == nil {
		return
	}
	PostOrderTree(root.Left)
	PostOrderTree(root.Right)
	fmt.Printf("tree node: %v", root.Val)
}

func Insert[T constraints.Ordered](root *Node[T], val T) *Node[T] {
	if root == nil {
		return &Node[T]{Val: val}
	}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
		return root
	}
	root.Right = Insert(root.Right, val)
	return root
}

func Insert1[T constraints.Ordered](root *Node[T], val T) *Node[T] {
	if root == nil {
		return &Node[T]{Val: val}
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

func CreateBiTree() *Node[int] {
	var input int
	var root *Node[int]
	for {
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			break
		}
		root = Insert1[int](root, input)
	}
	return root
}
