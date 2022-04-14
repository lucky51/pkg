package list

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrNotInitialized = errors.New("slist is not initialized")
	ErrNotFound       = errors.New("not found")
	ErrOutOfRange     = errors.New("out of range")
)

// ListNode ss
type ListNode[T any] struct {
	Data T
	Next *ListNode[T]
}

func (l *ListNode[T]) String() string {
	return fmt.Sprintf("%d", l.Data)
}

type SlinkedList[T any] struct {
	size int
	head *ListNode[T]
	tail *ListNode[T]
}

// InitSlinkedList 初始化链表
func InitSlinkedList[T any](s *SlinkedList[T]) {
	var result T
	s.head = &ListNode[T]{
		Data: result,
	}
	s.tail = s.head
}

func createSLinkedList[T any](s *SlinkedList[T], Data ...T) (*ListNode[T], *ListNode[T]) {
	var head, current *ListNode[T]
	for _, item := range Data {
		if current == nil {
			current = &ListNode[T]{
				Data: item,
			}
			head = current
		} else {
			current.Next = &ListNode[T]{Data: item}
		}
		current = current.Next
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
		for current.Next != nil {
			fmt.Fprintf(writer, "N:%d ->", current.Data)
			current = current.Next
		}
		fmt.Fprintf(writer, "N:%d", current.Data)
		fmt.Fprint(writer, "<-tail \n")
	}
}

func (s *SlinkedList[T]) Size() int { return s.size }
func (s *SlinkedList[T]) Append(Data T) (*ListNode[T], error) {
	if s == nil || s.head == nil {
		return nil, ErrNotInitialized
	}
	s.tail.Next = &ListNode[T]{
		Data: Data,
	}
	s.tail = s.tail.Next
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

// EachWithBreak 循环链表 f:返回false可以提前终止循环
func (s *SlinkedList[T]) EachWithBreak(f func(*ListNode[T]) bool) error {
	if s == nil {
		return ErrNotInitialized
	}
	current := s.head
	for current != nil {
		if !f(current) {
			return nil
		}
		current = current.Next
	}
	return nil

}

// Get 根据索引获取节点元素
func (s *SlinkedList[T]) Get(index int) (*ListNode[T], error) {
	if s == nil {
		return nil, ErrNotInitialized
	}
	if s.size < 1 {
		return nil, ErrNotFound
	}
	if index >= int(s.size) {
		return nil, ErrOutOfRange
	}
	if index == 0 {
		return s.head, nil
	}
	current := s.head
	var i int = 0
	for ; i <= index && current != nil; i++ {
		if current.Next != nil {
			current = current.Next
		}
	}
	if current == nil && i != index {
		return nil, ErrNotFound
	}
	return current, nil
}

// Insert 插入节点
func (s *SlinkedList[T]) Insert(index int, Data T) (*ListNode[T], error) {
	if s == nil {
		return nil, ErrNotInitialized
	}
	if index < 0 || index >= int(s.size) {
		return nil, ErrOutOfRange
	}
	newNode := &ListNode[T]{
		Data,
		nil,
	}
	if index-1 < 1 {
		index = 0
	}
	n, err := s.Get(index)

	if err != nil {
		return nil, err
	}
	temp := n.Next
	n.Next = newNode
	newNode.Next = temp
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
	if index < 0 || index > int(s.size-1) {

	}
	var deletedItem *ListNode[T]
	prevIndex := index - 1
	if prevIndex < 0 {
		prevIndex = 0
	}
	if prevIndex == index {
		deletedItem = s.head
		s.head = nil
		s.tail = nil
		s.size--
	} else {
		n, err := s.Get(prevIndex)
		if err != nil {
			return nil, err
		}
		deletedItem = n.Next
		n.Next = n.Next.Next
		s.size--
		if prevIndex == s.size-1 {
			s.tail = n
		}
	}
	return deletedItem, nil
}

// NewSLinkedList 创建链表
func NewSLinkedList[T any](Data ...T) *SlinkedList[T] {
	var s *SlinkedList[T] = new(SlinkedList[T])
	InitSlinkedList(s)
	s.head, s.tail = createSLinkedList(s, Data...)
	s.size = len(Data)
	return s
}
