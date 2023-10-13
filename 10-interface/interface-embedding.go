package main

import "io"

type student struct {
}

func (s student) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s student) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	var s student
	//The io.Reader interface in Go signifies the contract for types that implement a Read method.
	//If the student type has a Read method matching the required signature, this assignment will work.
	var r io.Reader = s
	var w io.Writer = s

	//any type that implements both Read and Write methods matches io.ReadWriter interface. If student has both Read and Write method, this assignment is valid.
	var rw io.ReadWriter = s
	_, _, _ = r, w, rw
}
