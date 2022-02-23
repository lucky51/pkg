package main

import (
	"fmt"
	"os"

	stack "github.com/lucky51/pkg/stack"
)

func main() {
	// s := slist.NewSlinkedList(1, 2, 3, 4, 5, 6, 7, 8)
	// s.Print(os.Stdout)
	// s.Insert(int(s.Size()-1), 9)
	// s.Insert(int(s.Size()-2), 10)
	// s.Print(os.Stdout)
	// s.Delete(1)
	// s.Print(os.Stdout)
	// s.Delete(1)
	// s.Print(os.Stdout)
	// s.Delete(1)
	// s.Print(os.Stdout)
	// s.Delete(1)
	// s.Print(os.Stdout)
	// s.Delete(1)
	// s.Print(os.Stdout)

	s1 := stack.NewStack()
	n, err := s1.Pop()
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("n=%v\n", n)
	}

	s1.Push(1)
	s1.Push(2)
	s1.Push(3)
	s1.Push(4)
	s1.Push(5)
	s1.Print(os.Stdout)
	s1.Pop()
	s1.Print(os.Stdout)
	s1.Pop()
	s1.Print(os.Stdout)
	s1.Pop()
	s1.Print(os.Stdout)
	s1.Pop()
	s1.Print(os.Stdout)
	n2, err := s1.Peek()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v", n2)
	s1.Print(os.Stdout)
}
