package main

import (
	"fmt"
)

func main() {
	linkedlist := NewList()

	for i := 0; i < 5; i++ {
		linkedlist.Add(i)
	}

	showInfoList(linkedlist)

	linkedlist.Pop()
	fmt.Println("After delete last element")
	showInfoList(linkedlist)

	val, err := linkedlist.At(5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(val)
	}

	linkedlist.DeleteFrom(3)

	fmt.Println("After delete operation: ")
	showInfoList(linkedlist)

	linkedlist.UpdateAt(2, 12)
	fmt.Println("After updating at pos: ")
	showInfoList(linkedlist)
}

func showInfoList(l *LinkedList) {
	fmt.Printf("Size = [%d]\n", l.Size)
	for j := l.Head; j != nil; j = j.Next {
		fmt.Printf("%v\n", j)
	}
}
