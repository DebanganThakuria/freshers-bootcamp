package main

import "fmt"

type tree struct {
	data byte
	left, right *tree
}

func buildTree(ex string) *tree {
	//fmt.Println("String inside buildTree", ex)
	if len(ex) == 1 {
		node := tree{ex[0], nil, nil}
		return &node
	}

	n1 := tree{ex[0], nil, nil}

	n3 := buildTree(ex[2:])

	n2 := tree{ex[1], &n1, n3}

	return &n2
}

func preorder(root *tree) {
	if root != nil {
		fmt.Printf("%c", root.data)
		preorder(root.left)
		preorder(root.right)
	}
}

func postorder(root *tree) {
	if root != nil {
		postorder(root.left)
		postorder(root.right)
		fmt.Printf("%c", root.data)
	}
}

func main() {
	var expression = "a+b-c"
	var root *tree

	root = buildTree(expression)

	preorder(root)
	fmt.Println("\n")
	postorder(root)
}