package test

import (
	"fmt"
	"github.com/lucky51/pkg/list"
	"os"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := list.NewSLinkedList(1, 2, 3, 4)
	fmt.Println("size", l.Size())
	list.Print(l, os.Stdout)
	item, _ := l.Get(0)
	fmt.Println("get 0:", item)
	l.Insert(0, 5)
	list.Print(l, os.Stdout)
	l.Insert(0, 6)
	fmt.Println("size", l.Size())
	l.Insert(0, 7)
	l.Append(8)
	l.Append(9)
	fmt.Println("size", l.Size())

	list.Print(l, os.Stdout)
	item1, _ := l.Delete(1)
	fmt.Println("delete:", item1)
	list.Print(l, os.Stdout)
	item2, _ := l.Delete(3)
	fmt.Println("delete:", item2)
	list.Print(l, os.Stdout)
	item3, _ := l.Delete(l.Size() - 1)
	fmt.Println("delete:", item3)
	list.Print(l, os.Stdout)
}
