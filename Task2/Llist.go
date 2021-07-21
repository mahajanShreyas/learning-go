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
	logger *zap.Logger
}

func (ll *LinkedList) Init() {
	ll.logger = zap.NewExample()
	defer ll.logger.Sync()
}

func (ll *LinkedList) Insert(data int) {

	if ll.start == nil {
		ll.logger.Debug("Value inserting at root",)

		ll.start = &node{data:data}

	} else {

		ll.logger.Debug("Value inserting at end",)
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
	ll.logger.Info("Printed Inorder Traversal of BST",
			zap.String("statusCode", "success..."),)
	temp := ll.start
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}


