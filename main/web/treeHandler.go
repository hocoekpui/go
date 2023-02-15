package main

import (
	"errors"
	"net/http"
	"strings"
)

var ErrorInvalidRouterPattern = errors.New("invalid router pattern")

type HandlerBasedOnTree struct {
	root *TreeNode
}

func NewHandlerBasedOnTree() Handler {
	rootNode := &TreeNode{}
	return &HandlerBasedOnTree{root: rootNode}
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	handlerFunc, ok := h.findRouter(h.root, c.R.URL.Path)
	if ok {
		handlerFunc(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not found"))
	}
}

func (h *HandlerBasedOnTree) Route(method string, path string, handlerFunc handlerFunc) error {
	err := h.validatePattern(path)
	if err != nil {
		return err
	}
	currentNode := h.root
	paths := strings.Split(strings.Trim(path, "/"), "/")
	for index, currentPath := range paths {
		childNode, ok := h.findMatchChild(currentNode, currentPath)
		if ok {
			currentNode = childNode
		} else {
			h.createChildNode(currentNode, paths[index:], handlerFunc)
			return nil
		}
	}
	currentNode.handlerFunc = handlerFunc
	return nil
}

func (h *HandlerBasedOnTree) findMatchChild(node *TreeNode, path string) (*TreeNode, bool) {
	var wildCardNode *TreeNode
	for _, childrenNode := range node.children {
		if childrenNode.path == path && childrenNode.path != "*" {
			return childrenNode, true
		} else if childrenNode.path == "*" {
			wildCardNode = childrenNode
		}
	}
	return wildCardNode, wildCardNode != nil
}

func (h *HandlerBasedOnTree) createChildNode(node *TreeNode, paths []string, handlerFunc handlerFunc) {
	currentNode := node
	for _, path := range paths {
		newNode := NewTreeNode(path)
		currentNode.children = append(currentNode.children, newNode)
		currentNode = newNode
	}
	currentNode.handlerFunc = handlerFunc
}

func (h *HandlerBasedOnTree) findRouter(node *TreeNode, path string) (handlerFunc, bool) {
	currentNode := node
	paths := strings.Split(strings.Trim(path, "/"), "/")
	for _, currentPath := range paths {
		childNode, ok := h.findMatchChild(currentNode, currentPath)
		if !ok {
			return nil, false
		}
		currentNode = childNode
	}
	if currentNode.handlerFunc == nil {
		return nil, false
	}
	return currentNode.handlerFunc, true
}

func (h *HandlerBasedOnTree) validatePattern(pattern string) error {
	pos := strings.Index(pattern, "*")
	if pos > 0 {
		if pos != len(pattern)-1 {
			return ErrorInvalidRouterPattern
		}
		if pattern[pos-1] != '/' {
			return ErrorInvalidRouterPattern
		}
	}
	return nil
}
