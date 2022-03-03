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

type ListNode[T any] struct {
	data *T
	next *ListNode[T]
}

func GetValue[T any](l *ListNode[T]) *T {
	var result *T
	if l == nil {
		return result
	}
	return l.data

}

func (l *ListNode[T]) GetData() *T {
	var result *T
	if l == nil {
		return result
	}
	return l.data
}

func (l *ListNode[T]) String() string {
	return fmt.Sprintf("%d", l.data)
}

type SlinkedList[T any] struct {
	size uint
	head *ListNode[T]
	tail *ListNode[T]
}

// InitSlinkedList 初始化链表
func InitSlinkedList[T any](s *SlinkedList[T]) {
	var result *T
	s.head = &ListNode[T]{
		data: result,
	}
	s.tail = s.head
}

func createSLinkedlist[T any](s *SlinkedList[T], data ...*T) (*ListNode[T], *ListNode[T]) {
	head := &ListNode[T]{}
	current := head
	for _, item := range data {
		current.next = &ListNode[T]{data: item}
		current = current.next
	}
	return head, current
}

// Print 打印链表
func Print[T any](s *SlinkedList[T], writer io.Writer) {
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

func (s *SlinkedList[T]) Size() uint { return s.size }
func (s *SlinkedList[T]) Append(data *T) (*ListNode[T], error) {
	if s == nil || s.head == nil {
		return nil, ErrNotInitialized
	}
	s.tail.next = &ListNode[T]{
		data: data,
	}
	s.tail = s.tail.next
	s.size++
	return s.tail, nil
}

// Each 循环链表
func (s *SlinkedList[T]) Each(f func(*ListNode[T])) error {
	s.EachWithBreak(func(node *ListNode[T]) bool {
		f(node)
		return true
	})
	return nil
}

// Each 循环链表 f:返回false可以提前终止循环
func (s *SlinkedList[T]) EachWithBreak(f func(*ListNode[T]) bool) error {
	if s == nil || s.head == nil {
		return ErrNotInitialized
	}
	current := s.head.next
	for current != nil {
		if !f(current) {
			return nil
		}
		current = current.next
	}
	return nil

}

// Get 根据索引获取节点元素
func (s *SlinkedList[T]) Get(index int) (*ListNode[T], error) {
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
func (s *SlinkedList[T]) Insert(index int, data *T) (*ListNode[T], error) {
	if s == nil || s.head == nil {
		return nil, ErrNotInitialized
	}
	if index > 0 && index >= int(s.size) {
		return nil, ErrOutofRange
	}
	newNode := &ListNode[T]{
		data,
		nil,
	}
	var n *ListNode[T]
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
func (s *SlinkedList[T]) Delete(index int) (*ListNode[T], error) {
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

func NewSlinkedList[T any](data ...*T) *SlinkedList[T] {
	var s *SlinkedList[T] = new(SlinkedList[T])
	InitSlinkedList(s)
	s.head, s.tail = createSLinkedlist(s, data...)
	s.size = uint(len(data))
	return s
}
