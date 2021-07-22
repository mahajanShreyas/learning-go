package main

import (
	"fmt"
	"go.uber.org/zap"
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
		
		zap.S().Debug("Value inserted at root ", data)

		bst.root = &node{data:data}
	} else {
		zap.S().Info("Non root value")
		insertUtil(bst.root, data)
	}
}

func insertUtil(root *node, data int) {
	if root.data > data {
		if root.left == nil {

			zap.S().Debug("Position found, Left Child of: ", root.data)

			root.left = &node{data:data}
		} else {

			zap.S().Info("Finding Position, Left Child of: ", root.data)

			insertUtil(root.left, data)
		}
	} else {
		if root.right == nil {

			zap.S().Debug("Position found, Right Child of: ", root.data)

			root.right = &node{data:data}
		} else {

			zap.S().Info("Finding Position, Right Child of: ", root.data)

			insertUtil(root.right, data)
		}
	}
}

func (bst Bst) PreOrderTraversal(){

	zap.S().Info("Printing Inorder Traversal of BST")

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