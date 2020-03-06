package main

import (
	"fmt"
	"io"
)

type Value interface{}

type NamedValue struct {
	Name  string
	Value Value
}

type ErrorValue struct {
	NamedValue
	Error error
}

func main() {
	var err error = io.EOF
	e := ErrorValue{NamedValue: NamedValue{Name: "fine", Value: 33}, Error: err}
	fmt.Printf("%+v\n", e)
	e = ErrorValue{NamedValue{"fine", 33}, err}
	fmt.Printf("%+v\n", e)
}
