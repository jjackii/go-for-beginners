package main

import "fmt"


func main() {
	go_slice := []string{"wozl", "jjack", "jacky"}
	go_slice = append(go_slice, "gopher")
	fmt.Println(go_slice)
}