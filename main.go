package main

import "fmt"


func main() {
	gopher := map[string]string{"name": "gopher", "birth": "2009"}
	for _, value := range gopher {
		fmt.Println(value)
	}
}