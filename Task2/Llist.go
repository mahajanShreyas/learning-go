package main

import (
	"fmt"
)

type node struct {
	data  int
	next  *node
}

type LinkedList struct {
	start *node
}


func (ll *LinkedList) Insert(data int) {

	if ll.start == nil {
		ll.start = &node{data:data}
	} else {
		temp := ll.start
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &node{data:data}
	}
}

func (ll *LinkedList) Traverse() {
	temp := ll.start
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}


