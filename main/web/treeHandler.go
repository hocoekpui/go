package main

import (
	"errors"
	"net/http"
	"sort"
	"strings"
)

var supportMethods = [4]string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}

var ErrorInvalidMethod = errors.New("invalid method")
var ErrorInvalidRouterPattern = errors.New("invalid router pattern")

type HandlerBasedOnTree struct {
	forest map[string]*TreeNode
}

func NewHandlerBasedOnTree() Handler {
	forest := make(map[string]*TreeNode, len(supportMethods))
	for _, method := range supportMethods {
		forest[method] = NewRootNode(method)
	}
	return &HandlerBasedOnTree{forest: forest}
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	handlerFunc, ok := h.findRouter(c.R.Method, c.R.URL.Path, c)
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
	currentNode, ok := h.forest[method]
	if !ok {
		return ErrorInvalidMethod
	}
	paths := strings.Split(strings.Trim(path, "/"), "/")
	for index, currentPath := range paths {
		childNode, ok := h.findMatchChild(currentNode, currentPath, nil)
		if ok && childNode.nodeType != nodeTypeAny {
			currentNode = childNode
		} else {
			h.createChildNode(currentNode, paths[index:], handlerFunc)
			return nil
		}
	}
	currentNode.handlerFunc = handlerFunc
	return nil
}

func (h *HandlerBasedOnTree) findMatchChild(node *TreeNode, path string, c *Context) (*TreeNode, bool) {
	candidates := make([]*TreeNode, 0, 2)
	for _, childrenNode := range node.children {
		if childrenNode.matchFunc(path, c) {
			candidates = append(candidates, childrenNode)
		}
	}
	if len(candidates) == 0 {
		return nil, false
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].nodeType < candidates[j].nodeType
	})
	return candidates[len(candidates)-1], true
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

func (h *HandlerBasedOnTree) findRouter(method string, path string, c *Context) (handlerFunc, bool) {
	currentNode, ok := h.forest[method]
	if !ok {
		return nil, false
	}
	paths := strings.Split(strings.Trim(path, "/"), "/")
	for _, currentPath := range paths {
		childNode, ok := h.findMatchChild(currentNode, currentPath, c)
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
