package queue

import (
	"fmt"
	"testing"

	"github.com/lucky51/pkg/tree"
)

func TestQueue(t *testing.T) {
	// q := NewQueue[int]()
	// q.Offer(1)
	// q.Offer(2)
	// q.Offer(3)
	// q.Offer(5)

	// slist.Print(q.sList, os.Stdout)

	// fmt.Println(q.Peek().GetData())
	// fmt.Println(q.Poll().GetData())
	// fmt.Println(q.Poll().GetData())
	// fmt.Println(q.Poll().GetData())
	// fmt.Println(q.Poll().GetData())
	// fmt.Println(q.Poll().GetData())
	// slist.Print(q.sList, os.Stdout)

	q1 := NewQueue[tree.TreeNode]()
	q1.Offer(&tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 4,
			},
		},
	})
	q1.Offer(&tree.TreeNode{
		Val: 5,
	})
	tree.MiddleOrderTree(q1.Peek())
	fmt.Println("queue size :", q1.Size())
	q1.Each(func(t *tree.TreeNode) {
		fmt.Println("each:", t.Val)
	})
	tree.MiddleOrderTree(q1.Poll())
	fmt.Println("queue size :", q1.Size())
	tree.MiddleOrderTree(q1.Poll())
	fmt.Println("queue size :", q1.Size())
	tree.MiddleOrderTree(q1.Poll())
	fmt.Println("queue size :", q1.Size())
	q1.Offer(&tree.TreeNode{
		Val: 10,
	})
	fmt.Println("queue size :", q1.Size())

}
