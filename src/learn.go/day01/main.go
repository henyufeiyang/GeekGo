package main

import "fmt"

func main() {
	/*
		fmt.Println("             *           ")
		fmt.Println("           *****         ")
		fmt.Println("          *******        ")
		fmt.Println("   ********************* ")
		fmt.Println("      ***************    ")
		fmt.Println("       *************     ")
		fmt.Println("      ****       ****    ")
		fmt.Println("     ***           ***   ")
		fmt.Println("    *                 *  ")

	*/

	a := []int{}
	fmt.Println(a)
	b := []int{1, 2, 3}
	fmt.Println("添加元素")
	b = append(b, 4)
	b = append(b, 5, 6)
	fmt.Println(b)
	fmt.Println("删除元素")
	b = append(b[:2], b[5:]...)
	fmt.Println(b)

	var c string = "hello"
	fmt.Println(c)
	cBytes := []byte(c)
	fmt.Println(cBytes)
	fmt.Println("修改切片内容")
	cBytes[0] = 'H'
	c = string(cBytes)
	fmt.Println(c)
}
