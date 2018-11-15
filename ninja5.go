package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		last      string
		favflavor []string
	}{
		first:     "Jane",
		last:      "Cane",
		favflavor: []string{"Vanila", "Chocolate"},
	}

	fmt.Println(p1)
	for i, v := range p1.favflavor {
		fmt.Println(i, v)
	}

}
