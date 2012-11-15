package main

import "fmt"

func plusX(x int)(func (int)(int)){
	return func (y int)(int){ return x*y} 
}

func main(){
	a := plusX(2)
	fmt.Printf("%d\n", a(3))
}

