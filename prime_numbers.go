package main

import "fmt"

func main(){
	var x int = 600851475143
	if x == 1{
		return false;
	}
	for i := 2; i*i <= x; i++{
		if x % i == 0{
			return false;
		}
	}
	fmt.Println(i)
}