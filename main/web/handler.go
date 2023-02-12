package main

import (
	"errors"
	"net/http"
	"strings"
)

var ErrorInvalidRouterPattern = errors.New("invalid router pattern")

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  handleFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	handler, ok := h.findRouter(c.R.URL.Path)
	if !ok {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h *HandlerBasedOnTree) Route(method string, pattern string, handlerFunc handleFunc) error {

	err := h.validatePattern(pattern)
	if err != nil {
		return err
	}

	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := h.root
	for index, path := range paths {
		matchChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = matchChild
		} else {
			h.createSubTree(cur, paths[index:], handlerFunc)
			return nil
		}
	}
	cur.handler = handlerFunc
	return nil
}

func (h *HandlerBasedOnTree) findMatchChild(root *node, path string) (*node, bool) {
	var wildcardNode *node
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
		if child.path == "*" {
			wildcardNode = child
		}
	}
	return wildcardNode, wildcardNode != nil
}

func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handleFunc handleFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handleFunc
}

func (h *HandlerBasedOnTree) findRouter(path string) (handleFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		child, ok := h.findMatchChild(cur, p)
		if !ok {
			return nil, false
		}
		cur = child
	}
	if cur == nil {
		return nil, false
	}
	return cur.handler, true
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

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2)}
}

var _ Handler = &HandlerBasedOnTree{}

func NewHandlerBasedOnTree() Handler {
	return &HandlerBasedOnTree{root: newNode("/")}
}

type Handler interface {
	ServeHTTP(c *Context)
	Routable
}
