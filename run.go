package main

import "fmt"

type rankholder interface {
	up()
}

type student struct {
	name string
	rank int
}

func (s student) String() string {
	return fmt.Sprintf("%s is the name \n%d is the rank", s.name, s.rank)
}

func (s *student) up() {
	s.rank++
}

func main() {
	var test rankholder = &student{"kavin", 123}
	test.up()
	fmt.Println(test.(*student))
}
