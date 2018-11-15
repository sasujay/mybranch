package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7,8}
	tot := tot(xi...)
	fmt.Println(tot)
}

func tot(x ...int) int {
	sum := 0
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for _, v := range x {
		sum += v
	}
	return sum
}
