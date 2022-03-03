package stack

import (
	"io"
	"os"

	slist "github.com/lucky51/pkg/list"
)

type Stack[T any] struct {
	sList *slist.SlinkedList[T]
	size  int
}

func (s *Stack[T]) Size() uint {
	return s.sList.Size()
}
func (s *Stack[T]) IsEmpty() bool {
	return s.sList.Size() == 0
}

// Pop 弹出栈顶元素
func (s *Stack[T]) Pop() (*slist.ListNode[T], error) {
	return s.sList.Delete(0)
}

// Push添加元素
func (s *Stack[T]) Push(data *T) (*slist.ListNode[T], error) {
	return s.sList.Insert(0, data)
}

// Peek 获取栈顶元素
func (s *Stack[T]) Peek() (*slist.ListNode[T], error) {
	return s.sList.Get(0)
}

// NewStack 创建栈
func NewStack[T any]() *Stack[T] {
	var s *slist.SlinkedList[T] = new(slist.SlinkedList[T])
	slist.InitSlinkedList(s)
	return &Stack[T]{
		sList: new(slist.SlinkedList[T]),
	}
}

// Print 打印
func (s *Stack[T]) Print(w io.Writer) {
	slist.Print(s.sList, os.Stdout)
}
