package queue

import (
	"fmt"
	"os"
	"testing"

	slist "github.com/lucky51/pkg/list"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	q.Offer(1)
	q.Offer(2)
	q.Offer(3)
	q.Offer(5)

	slist.Print(q.sList, os.Stdout)

	fmt.Println(q.Peek().GetData())
	fmt.Println(q.Poll().GetData())
	fmt.Println(q.Poll().GetData())
	fmt.Println(q.Poll().GetData())
	fmt.Println(q.Poll().GetData())
	fmt.Println(q.Poll().GetData())
	slist.Print(q.sList, os.Stdout)

}
