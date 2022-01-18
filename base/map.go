package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	m1["Tom"] = 18

	fmt.Printf(m2["Sam"])
}
