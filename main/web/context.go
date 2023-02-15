package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	PathParams map[string]string
}

func (c *Context) ReadJson(data interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func (c *Context) WriteJson(code int, data interface{}) error {
	response, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(response)
	if err != nil {
		return err
	}
	c.W.WriteHeader(code)
	return nil
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r, PathParams: make(map[string]string, 1)}
}
