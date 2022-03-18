package main

import (
	"errors"
	"fmt"
)

func alert(err error) {
	fmt.Println(err.Error())
}

type MyError struct {
	code    int
	message string
}

func (m *MyError) Error() string {
	return fmt.Sprintf("error %d: %s", m.code, m.message)
}

func main() {
	alert(errors.New("example error"))
	alert(&MyError{code: 404, message: "page not found"})
}
