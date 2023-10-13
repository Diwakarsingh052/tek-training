package main

//
//import (
//	"fmt"
//)
//
////When a method is defined with a value receiver,
////and a pointer calls that method, Go automatically dereferences the pointer to the underlying value
////and then calls the method.
//
//type Speaker interface {
//	Speak() string
//}
//
//type Person struct {
//	name string
//}
//
//func (p *Person) Speak() string {
//	return fmt.Sprintf("Hi, my name is %s", p.name)
//}
//
//func main() {
//	p := Person{}
//	p.Speak()
//	var s1 Speaker = &Person{"John"} // This is valid
//	fmt.Println(s1.Speak())
//
//	//var s2 Speaker = Person{"Jane"} // This will result in a compile error because Speak is a pointer receiver method
//	//fmt.Println(s2.Speak())
//}
//
////http.ListenAndServe(addr string, handler Handler) error
//
////The method set for a value, only includes methods implemented with a value receiver.
////The method set for a pointer, includes methods implemented with both pointer and value receivers.
////Methods declared with a pointer receiver, only implement the interface with pointer values.
////Methods declared with a value receiver, implement the interface with both a value and pointer receiver.
////The rules of method sets apply to interface types.
////Interfaces are reference types, don't share with a pointer.
////This is how we create polymorphic behavior in go.
