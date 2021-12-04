package main

import (
	"fmt"
)

func main() {
	// 计算两条直线是否平行
	a := [2][2]float64{}
	b := [2][2]float64{}
	fmt.Println("注意：输入点坐标时首先输入横坐标，回车后输入纵坐标")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			var x float64
			fmt.Printf("请分别输入第一条直线的第%d点坐标:", i+1)
			fmt.Scanln(&x)
			a[i][j] = x
			fmt.Println(a, b)
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			var y float64
			fmt.Printf("请分别输入第二条直线的第%d点坐标:", i+1)
			fmt.Scanln(&y)
			b[i][j] = y
			fmt.Println(a, b)
		}
	}

	ax := a[0][0] - a[1][0]
	ay := a[0][1] - a[1][1]
	bx := b[0][0] - b[1][0]
	by := b[0][1] - b[1][1]
	if ax == 0 && bx == 0 {
		fmt.Println("第一条直线平行于第二条直线")
	} else if by == 0 && ay == 0 {
		fmt.Println("第一条直线平行于第二条直线")
	} else if by != 0 && ay != 0 && ax != 0 && bx != 0 && ax/ay == bx/by {
		// 考虑除数和被除数为0的情况
		fmt.Println("第一条直线平行于第二条直线")
	} else {
		fmt.Println("两条直线不平行")
	}
	//fmt.Println("a: ", a)
	//fmt.Println("b: ", b)
}
