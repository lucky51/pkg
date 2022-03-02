package queue

import (
	"fmt"
	"os"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Offer(1)
	q.Offer(2)
	q.Offer(3)
	q.Offer(5)
	q.sList.Print(os.Stdout)
	fmt.Println(q.Peek().GetData())
	fmt.Println(q.Poll().GetData())
	fmt.Println(q.Poll().GetData())
	fmt.Println(q.Poll().GetData())
	fmt.Println(q.Poll().GetData())
	fmt.Println(q.Poll().GetData())
	q.sList.Print(os.Stdout)

}
