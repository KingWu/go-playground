package model

import (
	"fmt"
)

type Employee struct {
    ID          int32
    FirstName   string
    LastName    string
    BadgeNumber int32
}

// Demo2 print hello
func Demo2()  {
    fmt.Println("hello model")
}