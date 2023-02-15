package main

type TreeNode struct {
	path        string
	children    []*TreeNode
	handlerFunc handlerFunc
}

func NewTreeNode(path string) *TreeNode {
	return &TreeNode{path: path, children: make([]*TreeNode, 0, 2)}
}
