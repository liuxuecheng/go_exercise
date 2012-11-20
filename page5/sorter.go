/* Sort of multi-type interface */
package main

import (
	"fmt"
	)

type Sorter interface{
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Xi []int
type Xs []string

func (p Xi) Len() int{
	return len(p)
}

func (p Xi) Less(i,j int) bool{
	return p[j] < p[i]
}

func (p Xi) Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}

func (p Xs) Len() int{
	return len(p)
}

func (p Xs) Less(i,j int) bool{
	return p[j] < p[i]
}

func (p Xs) Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}

func Sort(x Sorter){
	for i := 0; i < x.Len() -1 ; i++{
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j){
				x.Swap(i,j)
			}
		}
	}
}

func main(){
	ints := Xi{44,12,34,-4,-7,3,199,200}
	strings := Xs{"go","hello","a","b","c","golang"}
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)
}
