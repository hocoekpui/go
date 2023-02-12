package main

import (
	"fmt"
	"time"
)

type handleFunc func(c *Context)

type FilterBuilder func(next Filter) Filter

type Filter func(c *Context)

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("Run time %d", end-start)
	}
}
