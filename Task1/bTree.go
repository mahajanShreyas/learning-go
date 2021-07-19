package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

type Bst struct {
	root *node
}

func (bst *Bst) Insert(data int) {

	if bst.root == nil {
		bst.root = &node{data:data}
	} else {
		insertUtil(bst.root, data)
	}
}

func insertUtil(root *node, data int) {
	if root.data > data {
		if root.left == nil {
			root.left = &node{data:data}
		} else {
			insertUtil(root.left, data)
		}
	} else {
		if root.right == nil {
			root.right = &node{data:data}
		} else {
			insertUtil(root.right, data)
		}
	}
}

func (bst Bst) PreOrderTraversal(){
	preOrderTraversalUtil(bst.root)
}

func preOrderTraversalUtil(root *node) {

	if root == nil {
		return
	}
	
	preOrderTraversalUtil(root.left)
	fmt.Println(root.data)
	preOrderTraversalUtil(root.right)
}