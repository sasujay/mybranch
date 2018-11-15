package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}
type barAttender struct {
	bar     string
	working bool
}

func (s secretAgent) speak() {
	fmt.Println("I am SA", s.first, s.last, " - in the secretAgent speak func")
}

func (p person) speak() {
	fmt.Println("I am person", p.first, p.last, " - in the person speak func")
}
func (b barAttender) speak() {
	fmt.Println("I am barAttender", b.bar, b.working, " - in the barAttender speak func")
}

type human interface {
	speak()
}

// Independent func used to call the polymorphed bar func
func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into func bar as person", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into func bar as SA", h.(secretAgent).first, h.(secretAgent).last)

	case barAttender:
		fmt.Println("I was passed into func bar as barAttender", h.(barAttender).bar)
	}
	fmt.Println("I was passed into func bar as human", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	b1 := barAttender{
		bar:     "SFO",
		working: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
	bar(b1)

}
