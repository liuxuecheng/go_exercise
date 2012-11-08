package main

import "fmt"

type stack struct{
	i int
	data [10]int
}

func (s *stack) push(k int){
	if s.i+1 >9{
		return 
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int{
	s.i-- 
	return s.data[s.i]
}

func main(){
	s := new(stack)
	s.push(18)
	s.push(23)
	fmt.Printf("Stack %v",s)
}
