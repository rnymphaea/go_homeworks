package main

import "errors"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewList() *LinkedList {
	n := &Node{}
	l := &LinkedList{
		Head: n}
	return l
}

func (l *LinkedList) Add(val int) {

	curr := l.Head
	if l.Size == 0 {
		curr.Val = val
	} else {
		for i := 0; i < l.Size-1; i++ {
			curr = curr.Next
		}
		next := &Node{
			Val: val,
		}
		curr.Next = next
	}
	l.Size++
}

func (l *LinkedList) Pop() {
	curr := l.Head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
	l.Size--
}

func (l *LinkedList) At(pos int) (int, error) {
	if pos > l.Size-1 {
		err := errors.New("Incorrect position")
		return 0, err
	}
	curr := l.Head
	for i := 0; i < pos; i++ {
		curr = curr.Next
	}
	return curr.Val, nil
}

func (l *LinkedList) DeleteFrom(pos int) {
	curr := l.Head
	if pos == 0 {
		l.Head = curr.Next
	} else if pos == l.Size-1 {
		l.Pop()
		return
	} else {
		for i := 0; i < pos-1; i++ {
			curr = curr.Next
		}
		next := curr.Next.Next
		curr.Next = next
	}
	l.Size--
}

func (l *LinkedList) UpdateAt(pos, val int) {
	curr := l.Head
	for i := 0; i < pos; i++ {
		curr = curr.Next
	}
	curr.Val = val
}
