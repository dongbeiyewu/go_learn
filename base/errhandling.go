package main

import (
	"errors"
	"fmt"
	"os"
)

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	fmt.Println("Hello,", name)
	return nil
}
func main() {
	_, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println(err)
	}
	if err := hello(""); err != nil {
		fmt.Println(err)
	}
}
