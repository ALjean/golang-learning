package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	err error
}

func (customError *CustomError) Error() string {
	return customError.err.Error()
}

func foo(num int) (res int, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = &CustomError{err: errors.New("my custom error")}
		}

	}()

	if num == 0 {
		panic("ops something wrong")
	}

	return num, err
}

func main() {

	res1, _ := foo(10)
	fmt.Printf("Result #%v\n", res1)
	_, error := foo(0)
	fmt.Printf("Error #%v\n", error)
}
