package list

import (
	"fmt"
	"os"
	"testing"
)

func TestLinkedList(t *testing.T) {
	//l := list.NewSLinkedList(1, 2, 3, 4)
	l := new(SLinkedList[int])
	fmt.Println("size", l.Size())
	//list.Print(l, os.Stdout)
	item, _ := l.Get(0)
	fmt.Println("get 0:", item)
	l.Insert(0, 5)
	Print(l, os.Stdout)
	l.Insert(0, 6)
	Print(l, os.Stdout)
	l.Insert(0, 7)
	Print(l, os.Stdout)
	//l.Append(8)
	//list.Print(l, os.Stdout)
	//l.Append(9)
	//list.Print(l, os.Stdout)
	//fmt.Println("size", l.Size())
	//item1, _ := l.Delete(1)
	//fmt.Println("delete:", item1)
	//list.Print(l, os.Stdout)
	//item2, _ := l.Delete(3)
	//fmt.Println("delete:", item2)
	//list.Print(l, os.Stdout)
	//item3, _ := l.Delete(l.Size() - 1)
	//fmt.Println("delete:", item3)
	//list.Print(l, os.Stdout)
}
