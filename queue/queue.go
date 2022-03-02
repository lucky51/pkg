package queue

import (
	"fmt"

	slist "github.com/lucky51/pkg/list"
)

type Queue[T any] struct {
	sList *slist.SlinkedList[T]
}

// offer 添加一个元素并返回true ，如果队列已满则返回false
func (q *Queue[T]) Offer(data T) bool {
	q.sList.Append(data)
	return true
}

// poll 移除并返回队列头部元素，如果队列为空，则返回nil
func (q *Queue[T]) Poll() *slist.ListNode[T] {
	if q.sList.Size() == 0 {
		return nil
	}
	node, err := q.sList.Delete(0)
	if err != nil {
		return nil
	}
	return node
}

// peek 返回队列头部元素，如果队列为空则返回nil
func (q *Queue[T]) Peek() *slist.ListNode[T] {
	n, err := q.sList.Get(0)
	if err == nil {
		return n
	}
	fmt.Println(err)
	return nil
}

// NewQueue 创建队列
func NewQueue[T any]() *Queue[T] {
	var s *slist.SlinkedList[T] = new(slist.SlinkedList[T])
	slist.InitSlinkedList(s)
	return &Queue[T]{sList: s}
}
