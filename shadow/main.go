package main

import "fmt"

func main() {
	shadowing()
}

func shadowing() {
	var x int = 1

	{
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
}
