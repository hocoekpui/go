package main

import (
	"fmt"
	"time"
)

type Filter func(c *Context)

type FilterBuilder func(filter Filter) Filter

func MetricFilterBuilder(filter Filter) Filter {
	return func(c *Context) {
		startTime := time.Now().UnixNano()
		filter(c)
		endTime := time.Now().UnixNano()
		fmt.Printf("Cost time: %d \n", endTime-startTime)
	}
}
