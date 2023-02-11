package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c Context) ReadJson(data interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func (c Context) WriteJson(status int, data interface{}) error {
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(result)
	if err != nil {
		return err
	}
	c.W.WriteHeader(status)
	return nil
}

func (c Context) Success(data interface{}) error {
	return c.WriteJson(http.StatusOK, data)
}

func (c Context) Fail(data interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, data)
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r}
}
