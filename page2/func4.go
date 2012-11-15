package main

import "fmt"

func myfun(arg ... int){
	for _, n := range arg{
		fmt.Printf("the number is : %d\n" ,n)
	}
}

func main(){
	myfun(1,4,5,6,7)
	myfun(1,6,3,4,6,1)
}
