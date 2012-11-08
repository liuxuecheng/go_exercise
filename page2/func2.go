package main

func order(a,b int)(a,b int){
	if a > b {
		return b,a
	}
	return a,b
}
