package main

import "fmt"

type Speaker interface {
	Speak() string
	Speak2() string
}

type Person struct {
	name string
}

func (p *Person) Speak() string {
	return fmt.Sprintf("Hi, my name is %s", p.name)
}
func (p Person) Speak2() string {
	return fmt.Sprintf("Hi, my name is 2 %s", p.name)
}

func main() {
	//p := Person{}
	//p.Speak()
	var s1 Speaker = &Person{"John"} // This is valid
	fmt.Println(s1.Speak())
	fmt.Println(s1.Speak2())
}
