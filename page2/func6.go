package main

import "fmt"

func Map(f func(int) int, l[]int)([]int){
	j := make([]int, len(l))
	for i, v:= range l{
		j[i] = f(v)
	}
	return j 
}

func square(i int)(value int){
	value = i*i
	return 
}

func main(){
	l := []int{1,2,3,4,5,6,7,8,9,10}
	f := square
	j := Map(f,l)
	fmt.Printf("v%",j)
}
