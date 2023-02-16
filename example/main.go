package main

import (
	"fmt"

	"github.com/jinyus/confirmop"
)

func main() {
	userChoice := confirmop.ConfirmOperation("Do you want to delete this file", "proceed", true, true)
	if userChoice {
		fmt.Println("Safe to delete the file")
	} else {
		fmt.Println("Not safe to delete the file")
	}
}
