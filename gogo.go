package main

import "fmt"
//для легкости действий функцию назвал фибоначчи
func fibonacci() func() int {
//начальные числа в фибоначчи всегда 0 и 1 
	x, y := 0, 1
	return func() int {
//принцип действия фибоначчи
		x, y = y, x + y
		return x
	}
}
//основаная функция что будет брать числа из необходимого фактора и вычислять их
func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}