package main

import (
	"errors"
	"fmt"
)

type BusinessError struct {
}

func (b BusinessError) Error() string {
	return "This is BusinessError"
}

func main() {
	err := &BusinessError{}
	wrappedErr := fmt.Errorf("this is a wrapped error %w", err)

	if err == errors.Unwrap(wrappedErr) {
		fmt.Println("error unwrapped")
	}

	if errors.Is(wrappedErr, err) {
		fmt.Println("wrapped is BusinessError")
	}

	/*类型转换*/
	newErr := &BusinessError{}
	errors.As(err, newErr)
}
