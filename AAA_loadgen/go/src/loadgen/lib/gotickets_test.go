package lib

import "fmt"

func ExampleSampleTickets_Take() {
	s, _ := NewSampleTickets(5)
	fmt.Println(s.Active())
	fmt.Println(s.Remainder())
	s.Take()
	fmt.Println(s.Remainder())
	//output: true
	//5
	//4
}

func ExampleSampleTickets_Take2() {
	s, _ := NewSampleTickets(5)
	fmt.Println(s.Active())
	fmt.Println(s.Remainder())
	s.Take()
	s.Take()
	fmt.Println(s.Remainder())
	s.Return()
	fmt.Println(s.Remainder())
	//output: true
	//5
	//3
	//4
}
