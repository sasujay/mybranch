package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, _ := file.Read(data)
	fmt.Printf("Read %d bytes: %q\n, cap %d\n", count, data[:count], cap(data))
}
