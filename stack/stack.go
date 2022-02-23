package stack

import (
	"io"

	slist "github.com/lucky51/pkg/list"
)

type Stack struct {
	sList *slist.SlinkedList
	size  int
}

func (s *Stack) Pop() (*slist.ListNode, error) {
	return s.sList.Delete(0)
}

func (s *Stack) Push(data int) (*slist.ListNode, error) {
	return s.sList.Insert(0, data)
}
func (s *Stack) Peek() (*slist.ListNode, error) {
	return s.sList.Get(0)
}
func NewStack() *Stack {
	var s *slist.SlinkedList = new(slist.SlinkedList)
	slist.InitSlinkedList(s)
	return &Stack{
		s,
		0,
	}
}
func (s *Stack) Print(w io.Writer) {
	s.sList.Print(w)
}
