package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	 vehicle
	fourwheels bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourwheels: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}
	fmt.Println(t1, s1)
	fmt.Println(t1.doors, s1.color)
}
