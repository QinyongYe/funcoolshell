package main

import (
	"strings"
	"fmt"
)

type BinaryTree struct {
	Left *BinaryTree
	Right *BinaryTree
	Value uint8
}

func BuildBinaryTree(inorder, postorder string) *BinaryTree {
	if len(inorder) != len(postorder) {
		panic("inorder not equal postorder")
	}

	tree := BinaryTree{Value:postorder[len(postorder)-1]}
	index := strings.IndexByte(inorder, byte(tree.Value))
	if index < 0 {
		panic(fmt.Sprint("no ", tree.Value, " in ", inorder))
	}

	if index > 0 { // has left tree
		leftin := inorder[:index]
		leftpost := postorder[:index]
		tree.Left = BuildBinaryTree(leftin, leftpost)
	}

	if index < len(inorder) - 1 { // has right tree
		rightin := inorder[index+1:]
		rightpost := postorder[index:len(inorder)-1]
		tree.Right = BuildBinaryTree(rightin, rightpost)
	}

	return &tree
}

func PrintBinaryTree(tree *BinaryTree) {
	printBinaryTree(tree, 0)
}

func printBinaryTree(tree *BinaryTree, indent int) {
	fmt.Println(getSpaces(indent), string(tree.Value))
	if tree.Left != nil {
		printBinaryTree(tree.Left, indent+1)
	}

	if tree.Right != nil {
		printBinaryTree(tree.Right, indent+1)
	}
}

func getSpaces(number int) string {
	result := ""
	for i := 0; i<number;i++ {
		result = result + " "
	}
	return result
}

