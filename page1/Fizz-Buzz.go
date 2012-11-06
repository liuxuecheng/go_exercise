package main

import "fmt"

func main(){
	for i := 1; i <= 100; i++ {
		switch{
			case i%3 == 0 && i%5 != 0:
				fmt.Println(i,"Fizz")
			case i%3 != 0 && i%5 == 0:
				fmt.Println(i,"Buzz")
			case i%3 ==0 && i%5 ==0:
				fmt.Println(i,"FizzBuzz")
			default:
				fmt.Println(i)
		}
	}
}
