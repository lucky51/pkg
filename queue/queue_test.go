package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 5; i++ {
		temp := i
		q.Offer(&temp)
	}
	q.Each(func(t *int) {
		fmt.Printf("%d,", *t)
	})
	fmt.Printf("\n peek last:\n")
	l := q.PeekLast()
	fmt.Println("peek last:", l)
	pl := q.PollLast()
	fmt.Println("poll last:", pl)
	q.Each(func(t *int) {
		fmt.Printf("%d,", *t)
	})
	fmt.Printf("\npoll last \n")
	q.Poll()
	q.Each(func(t *int) {
		fmt.Printf("%d,", *t)
	})
	fmt.Printf("\npoll  \n")
	// slist.Print(q.sList, os.Stdout)

	// fmt.Println(q.Peek().GetData())
	// fmt.Println(q.Poll().GetData())
	// fmt.Println(q.Poll().GetData())
	// fmt.Println(q.Poll().GetData())
	// fmt.Println(q.Poll().GetData())
	// fmt.Println(q.Poll().GetData())
	// slist.Print(q.sList, os.Stdout)

	// q1 := NewQueue[tree.TreeNode]()
	// q1.Offer(&tree.TreeNode{
	// 	Val: 1,
	// 	Left: &tree.TreeNode{
	// 		Val: 2,
	// 	},
	// 	Right: &tree.TreeNode{
	// 		Val: 3,
	// 		Left: &tree.TreeNode{
	// 			Val: 4,
	// 		},
	// 	},
	// })
	// q1.Offer(&tree.TreeNode{
	// 	Val: 5,
	// })
	// tree.MiddleOrderTree(q1.Peek())
	// fmt.Println("queue size :", q1.Size())
	// q1.Each(func(t *tree.TreeNode) {
	// 	fmt.Println("each:", t.Val)
	// })
	// tree.MiddleOrderTree(q1.Poll())
	// fmt.Println("queue size :", q1.Size())
	// tree.MiddleOrderTree(q1.Poll())
	// fmt.Println("queue size :", q1.Size())
	// tree.MiddleOrderTree(q1.Poll())
	// fmt.Println("queue size :", q1.Size())
	// q1.Offer(&tree.TreeNode{
	// 	Val: 10,
	// })
	// fmt.Println("queue size :", q1.Size())

}
