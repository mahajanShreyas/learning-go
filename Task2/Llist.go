package main

import (
	"fmt"
	"go.uber.org/zap"
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

		zap.L().Debug("Value inserting at root",)

		ll.start = &node{data:data}

	} else {

		zap.L().Debug("Value inserting at end",)

		temp := ll.start
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &node{data:data}
	}
}

func (ll *LinkedList) Traverse() {

	zap.L().Info("Printed Inorder Traversal of BST",
			zap.String("statusCode", "success..."),)

	temp := ll.start
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}


