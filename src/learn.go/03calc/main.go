package main

import "fmt"

func main() {
	/*
		var a, b int = 10, 3
		fmt.Println("a + b = ", a+b)
		fmt.Println("a - b = ", a-b)
		fmt.Println("a * b = ", a*b)
		fmt.Println("a / b = ", a/b)
		fmt.Println("a % b = ", a%b)
		fmt.Println(true && false == false)
		fmt.Println("a > b =", a > b)
		fmt.Println("a < b =", a < b)
	*/
	arr := []int{-1, 1, 2, 2, -3, -3, 4, 4, 5, 5, 6, 6, 7}
	result := -1
	fmt.Println(arr)
	for i, item := range arr {
		fmt.Println("i:", i, "item:", item)
		if result < 0 {
			result = item
		} else {
			result = result ^ item
		}
	}
	fmt.Println(result)
}
