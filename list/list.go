package list

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrNotInitialized = errors.New("slist is not initialized")
	ErrNotFound       = errors.New("not found")
	ErrOutofRange     = errors.New("out of range")
)

type ListNode struct {
	data int
	next *ListNode
}

func (l *ListNode) GetData() int {
	if l == nil {
		return -1
	}
	return l.data
}
func (l *ListNode) String() string {
	return fmt.Sprintf("%d", l.data)
}

type SlinkedList struct {
	size uint
	head *ListNode
	tail *ListNode
}

// InitSlinkedList 初始化链表
func InitSlinkedList(s *SlinkedList) {
	s.head = &ListNode{
		data: -1,
	}
	s.tail = s.head
}

func createSLinkedlist(s *SlinkedList, data ...int) (*ListNode, *ListNode) {
	head := &ListNode{data: -1}
	current := head
	for _, item := range data {
		current.next = &ListNode{data: item}
		current = current.next
	}
	return head, current
}

// Print 打印链表
func (s *SlinkedList) Print(writer io.Writer) {
	if s == nil || s.head == nil {
		fmt.Fprintf(writer, "null")
	} else {
		current := s.head
		fmt.Fprint(writer, "head->")
		for current.next != nil {
			fmt.Fprintf(writer, "N:%d ->", current.data)
			current = current.next
		}
		fmt.Fprintf(writer, "N:%d", current.data)
		fmt.Fprint(writer, "<-tail \n")
	}
}

func (s *SlinkedList) Size() uint { return s.size }
func (s *SlinkedList) Append(data int) (*ListNode, error) {
	if s == nil || s.head == nil {
		return nil, ErrNotInitialized
	}
	s.tail.next = &ListNode{
		data: data,
	}
	s.tail = s.tail.next
	s.size++
	return s.tail, nil
}

// Get 根据索引获取节点元素
func (s *SlinkedList) Get(index int) (*ListNode, error) {
	if s == nil || s.head == nil {
		return nil, ErrNotInitialized
	}
	if s.size < 1 {
		return nil, ErrNotFound
	}
	if index >= int(s.size) {
		return nil, ErrOutofRange
	}
	if index == 0 {
		return s.head.next, nil
	}
	current := s.head.next
	for i := 0; i < index; i++ {
		if current.next == nil {
			return nil, ErrNotFound
		} else {
			current = current.next
		}
	}
	return current, nil
}

// Insert 插入节点
func (s *SlinkedList) Insert(index int, data int) (*ListNode, error) {
	if s == nil || s.head == nil {
		return nil, ErrNotInitialized
	}
	if index > 0 && index >= int(s.size) {
		return nil, ErrOutofRange
	}
	newNode := &ListNode{
		data,
		nil,
	}
	var n *ListNode
	if index == 0 {
		n = s.head
	} else {
		var err error
		n, err = s.Get(index - 1)

		if err != nil {
			return nil, err
		}
	}
	temp := n.next
	n.next = newNode
	newNode.next = temp
	if index == int(s.size-1) {
		s.tail = newNode
	}
	s.size++
	return s.tail, nil
}

// Delete 删除链表元素
func (s *SlinkedList) Delete(index int) (*ListNode, error) {
	if s.size < 1 {
		return nil, ErrNotFound
	}
	if index == 0 {
		temp := s.head.next
		s.head.next = s.head.next.next
		if int(s.size-1) == 0 {
			s.tail = s.head
		}
		s.size--
		return temp, nil
	} else {
		n, err := s.Get(index - 1)
		if err != nil {
			return nil, err
		} else {
			deleting := n.next
			n.next = n.next.next
			if index == int(s.size-1) {
				s.tail = n
			}
			s.size--
			return deleting, nil
		}
	}
}

func NewSlinkedList(data ...int) *SlinkedList {
	var s *SlinkedList = new(SlinkedList)
	InitSlinkedList(s)
	s.head, s.tail = createSLinkedlist(s, data...)
	s.size = uint(len(data))
	return s
}
