package main

import "fmt"


func sum(x int, y int) int {
    return x + y
}

func main() {
	var sum int = sum(5,5)
	fmt.Printf("%d \n", sum)
}
