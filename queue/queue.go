package queue

import (
	"fmt"

	slist "github.com/lucky51/pkg/list"
)

type Queue[T any] struct {
	sList *slist.SlinkedList[T]
}

// Offer 添加一个元素并返回true ，如果队列已满则返回false
func (q *Queue[T]) Offer(data *T) bool {
	q.sList.Append(data)
	return true
}

// Size 返回队列大小
func (q *Queue[T]) Size() uint {
	return q.sList.Size()
}

// Each 循环队列元素
func (q *Queue[T]) Each(f func(*T)) error {
	return q.sList.Each(func(ln *slist.ListNode[T]) {
		f(ln.Data)
	})
}

// EachWithBreak 遍历队列 f: 如果返回false 提前终止循环
func (q *Queue[T]) EachWithBreak(f func(*slist.ListNode[T]) bool) error {
	return q.sList.EachWithBreak(func(ln *slist.ListNode[T]) bool {
		return f(ln)
	})
}

// IsEmpty 是否为空
func (q *Queue[T]) IsEmpty() bool {
	return q.sList.Size() == 0
}

// PollLast 移除并返回队尾元素
func (q *Queue[T]) PollLast() *T {
	if q.sList.Size() == 0 {
		return nil
	}
	node, err := q.sList.Delete(int(q.Size()) - 1)
	if err != nil {
		return nil
	}
	return node.Data
}

// Poll 移除并返回队列头部元素，如果队列为空，则返回nil
func (q *Queue[T]) Poll() *T {
	if q.sList.Size() == 0 {
		return nil
	}
	node, err := q.sList.Delete(0)
	if err != nil {
		return nil
	}
	return node.Data
}

// Peek 返回队列头部元素，如果队列为空则返回nil
func (q *Queue[T]) Peek() *T {
	n, err := q.sList.Get(0)
	if err == nil {
		return n.Data
	}
	fmt.Println(err)
	return nil
}

// PeekLast 返回队列尾部元素，如果队列为空则返回nil
func (q *Queue[T]) PeekLast() *T {
	n, err := q.sList.Get(int(q.Size()) - 1)
	if err != nil {
		return n.Data
	}
	return nil
}

// NewQueue 创建队列
func NewQueue[T any]() *Queue[T] {
	var s *slist.SlinkedList[T] = new(slist.SlinkedList[T])
	slist.InitSlinkedList(s)
	return &Queue[T]{sList: s}
}
