package main

import ("fmt")

func main() {
	var s interface{}
	b := 1
	s = b
	c := s.(int)	
	fmt.Println(c)
}
