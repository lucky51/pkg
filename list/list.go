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

// Node ss
type Node[T any] struct {
	Data T
	Next *Node[T]
}

func (l *Node[T]) String() string {
	return fmt.Sprintf("%d", l.Data)
}

type SLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// InitSlinkedList 初始化链表
func InitSlinkedList[T any](s *SLinkedList[T]) {
	var result T
	s.head = &Node[T]{
		Data: result,
	}
	s.tail = s.head
}

func createSLinkedList[T any](s *SLinkedList[T], Data ...T) (*Node[T], *Node[T]) {
	var head, current *Node[T]
	for _, item := range Data {
		if current == nil {
			current = &Node[T]{
				Data: item,
			}
			head = current
		} else {
			current.Next = &Node[T]{Data: item}
		}
		if current.Next != nil {
			current = current.Next
		}
	}
	return head, current
}

// Print 打印链表
func Print[T any](s *SLinkedList[T], writer io.Writer) {
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

func (s *SLinkedList[T]) Size() int { return s.size }
func (s *SLinkedList[T]) Append(Data T) (*Node[T], error) {
	if s == nil {
		return nil, ErrNotInitialized
	}
	s.tail.Next = &Node[T]{
		Data: Data,
	}
	s.tail = s.tail.Next
	s.size++
	return s.tail, nil
}

// Each 循环链表
func (s *SLinkedList[T]) Each(f func(*Node[T])) error {
	s.EachWithBreak(func(node *Node[T]) bool {
		f(node)
		return true
	})
	return nil
}

// EachWithBreak 循环链表 f:返回false可以提前终止循环
func (s *SLinkedList[T]) EachWithBreak(f func(*Node[T]) bool) error {
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
func (s *SLinkedList[T]) Get(index int) (*Node[T], error) {
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
	for ; i < index && current != nil; i++ {
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
func (s *SLinkedList[T]) Insert(index int, Data T) (*Node[T], error) {
	if s == nil {
		return nil, ErrNotInitialized
	}
	if index < 0 || index >= int(s.size) {
		return nil, ErrOutOfRange
	}
	newNode := &Node[T]{
		Data,
		nil,
	}
	prevIndex := index - 1
	if prevIndex < 1 {
		prevIndex = 0
	}
	n, err := s.Get(prevIndex)

	if err == ErrNotFound {
		s.head = newNode
		s.tail = s.head
		return s.head, nil
	}

	if prevIndex == 0 {
		newNode.Next = n
		s.head = newNode
	} else {
		temp := n.Next
		n.Next = newNode
		newNode.Next = temp
		if index == int(s.size-1) {
			s.tail = newNode
		}
	}
	s.size++
	return s.tail, nil
}

// Delete 删除链表元素
func (s *SLinkedList[T]) Delete(index int) (*Node[T], error) {
	if s.size < 1 {
		return nil, ErrNotFound
	}
	if index < 0 || index > int(s.size-1) {
		return nil, ErrOutOfRange
	}
	var deletedItem *Node[T]
	prevIndex := index - 1
	if prevIndex < 0 {
		prevIndex = 0
	}
	if prevIndex == index {
		deletedItem = s.head
		s.head = s.head.Next
		s.tail = nil
		s.size--
	} else {
		n, _ := s.Get(prevIndex)
		fmt.Println("delted get:", prevIndex, n)
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
func NewSLinkedList[T any](Data ...T) *SLinkedList[T] {
	var s *SLinkedList[T] = new(SLinkedList[T])
	InitSlinkedList(s)
	s.head, s.tail = createSLinkedList(s, Data...)
	s.size = len(Data)
	return s
}
