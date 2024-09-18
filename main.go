package main

import "fmt"

type mascot struct {
	name string
	bitrh int
	details []string
}

func main() {
	details := []string{"eyes", "teeth"}
	gopher := mascot {name: "go gopher", bitrh: 2009, details: details}
	fmt.Println(gopher.name)
}