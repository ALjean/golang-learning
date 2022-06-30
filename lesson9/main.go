package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	error error
}

func (customError *CustomError) Error() string {
	return customError.error.Error()
}

func NewCustomError(msg string) error {
	return &CustomError{
		error: errors.New(msg),
	}
}

var internalError = NewCustomError("internal error")

func GetClient() (string, error) {
	err := fmt.Errorf(" client error %w", internalError)
	return "result", err
}

func main() {

	res, err := GetClient()

	if err != nil {
		if errors.Is(err, internalError) {
			fmt.Println("Happened internal error")
		}

		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}
