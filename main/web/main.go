package main

import (
	"fmt"
	"net/http"
)

func index(c *Context) {
	_, _ = fmt.Fprintf(c.W, "index")
}

func match(c *Context) {
	_, _ = fmt.Fprintf(c.W, "match")
}

func matchDetail(c *Context) {
	_, _ = fmt.Fprintf(c.W, "match detail")
}

func main() {
	server := NewCustomerServer(MetricFilterBuilder)
	server.Route(http.MethodGet, "/index", index)
	server.Route(http.MethodGet, "/match/detail", matchDetail)
	server.Route(http.MethodGet, "/match/*", match)
	server.Start(":8080")
}
