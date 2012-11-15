package main

import "fmt"

func bubblesort(n []int){
	for i :=0; i< len(n)-1; i++{
		for j := i+1; j<len(n); j++{
			if n[j] < n[i]{
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}

func main(){
	l := []int{0,19,3,-4,-3,1,18,190,22,16}
	fmt.Printf("%v\n", l)
	bubblesort(l)
	fmt.Printf("%v\n", l)
}
