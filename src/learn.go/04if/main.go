package main

import (
	"fmt"
)

func main() {
	/*
		var fruit string = "6 apples"
		var watermelon bool = false
		if watermelon {
			fruit = "1 apples"
		}
		fmt.Println("buy: ", fruit)
	*/

	var rmb int = 200
	var busy bool = false
	/*if rmb <= 20 {
		fmt.Println("点外卖")
	} else if rmb <= 200 {
		fmt.Println("下馆子")
	} else if rmb <= 2000 {
		fmt.Println("去米其林餐厅")
	} else if rmb <= 20000 {
		fmt.Println("出去旅游一个月")
	} else {
		fmt.Println("买辆车")
	}
	*/
	switch {
	case rmb <= 20 && rmb >= 0:
		fmt.Println("点外卖")
		//fallthrough
	case rmb <= 200 && rmb > 20:
		fmt.Println("下馆子")
		if busy {
			break
		} else {
			fmt.Println("出去溜溜")
		}
	case rmb <= 2000 && rmb > 200:
		fmt.Println("去米其林")
	case rmb <= 20000 && rmb > 2000:
		fmt.Println("出国游")
	default:
		fmt.Println("容我想想")
	}
}
