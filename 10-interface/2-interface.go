package main

import "fmt"

// Polymorphism means that a piece of code changes its behavior depending on the
//concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike
// Bigger the interface weaker the abstraction // Rob Pike

type Speaker interface {
	Speak()
}
type human struct {
	name string
}

func (h human) Speak() {
	fmt.Println("human speaking", h.name)
}

type ai struct {
	name string
}

func (a ai) Speak() {
	fmt.Println("ai speaking", a.name)
}

// doSomething is a polymorphic func
// doSomething() will accept any type of value which implements reader interface

func doSomething(s Speaker) {
	fmt.Printf("%T\n", s)
	s.Speak()
}

func main() {
	h := human{name: "dev"}
	a := ai{name: "alexa"}
	//values of ai and human could be passed to doSomething because both struct implement the Speaker interface
	doSomething(h)
	doSomething(a)
}
