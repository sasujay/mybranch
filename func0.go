package main

import "fmt"

func main() {

	foo()
	bar("Jane")
	s1 := woo("Jack")
	fmt.Println(s1)
	x, y := numb("Hello", "Integeres")
	fmt.Println(x, y)
	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)

}

func foo() {
	fmt.Println("First Func")
}
func bar(s string) {
	fmt.Println("Hello", s)
}
func woo(s string) string {
	return fmt.Sprint("Hello ", s)
}

func numb(s string, s1 string) (string, bool) {
	a := fmt.Sprint(s, " ", s1, " and Floats")
	b := true
	return a, b
}
func sum(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
