package main
import "fmt"

type person struct {
	first string
	last string
}
type sAgent struct {
	person
	ltk bool
}
func (s sAgent) Roll() {
	fmt.Println("Awesome",s.first,s.last)
}
func (a person) Name() {
	fmt.Println("Cool Person",a.first,a.last)
}
func main() {
	sa1:= sAgent{
		person:person {
			first: "Jane",
			last: "Cane",
		},
		ltk: true,
	}

	p1:=person {
		first:"James",
		last:"Bond",
	}
fmt.Println(sa1)
	sa1.Roll()
	p1.Name()

}
