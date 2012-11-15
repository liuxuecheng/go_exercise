package main

import "fmt"

func max(l[]int)(m int){
	m = l[0]
	for _, v := range(l){
		if v > m{
			m = v
		}
	}
	return
}

func main(){
	l := []int{1,9,17,19,28,3,88,2,100,49,7,97}
	fmt.Printf("the max number is :%d\n", max(l))
}
