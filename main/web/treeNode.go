package main

import "strings"

const (
	nodeTypeRoot = iota

	nodeTypeAny

	nodeTypeParam

	nodeTypeStatic
)

type matchFunc func(path string, c *Context) bool

type TreeNode struct {
	path        string
	children    []*TreeNode
	handlerFunc handlerFunc
	matchFunc   matchFunc
	nodeType    int
}

func NewRootNode(method string) *TreeNode {
	return &TreeNode{
		path:     method,
		children: make([]*TreeNode, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			return false
		},
		nodeType: nodeTypeRoot,
	}
}

func NewStaticNode(path string) *TreeNode {
	return &TreeNode{
		path:     path,
		children: make([]*TreeNode, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			return path == p && p != "*"
		},
		nodeType: nodeTypeStatic,
	}
}

func NewAnyNode() *TreeNode {
	return &TreeNode{
		path: "*",
		matchFunc: func(p string, c *Context) bool {
			return true
		},
		nodeType: nodeTypeAny,
	}
}

func NewParamNode(path string) *TreeNode {
	param := path[1:]
	return &TreeNode{
		path:     path,
		children: make([]*TreeNode, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			if c != nil {
				c.PathParams[param] = p
			}
			return path != "*"
		},
		nodeType: nodeTypeParam,
	}
}

func NewTreeNode(path string) *TreeNode {
	if path == "*" {
		return NewAnyNode()
	}
	if strings.HasPrefix(path, ":") {
		return NewParamNode(path)
	}
	return NewStaticNode(path)
}
