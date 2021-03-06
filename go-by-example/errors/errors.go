package main

import (
    "errors"
    "fmt"
)
// In Go, you can return multiple return values. This allows you to return an explicit error value
// rather than a single result/error return value

// By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
    if arg == 42 {
        // errors.new constructs a basic error value with the supplied message
        return -1, errors.New("can't work with 42")

    }
    // A nil value in the error position indicates that there was no error
    return arg + 3, nil
}

type argError struct {
    arg  int
    prob string
}

// It’s possible to use custom types as errors by implementing the Error() method on them
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        // In this case we use &argError syntax to build a new struct,
        // supplying values for the two fields arg and prob.
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}