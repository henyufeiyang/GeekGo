package main

import "fmt"

func main() {
	/*
		var f1 float64 = 1.234
		var i1 int = int(f1)
		fmt.Println("f1:", f1, "i1:", i1)
		var i2 int = 22
		fmt.Println("int2:", i2)
		var hello string = "hello"
		fmt.Println("hello:", hello)
		f2 := 3.333
		fmt.Println("f2:", f2)
		var i3, i4 = 33, 44
		fmt.Println("i3:", i3, "i4:", i4)
		fmt.Println(i3 * i4)
		var he = i2 + i3
		fmt.Println("he:", he)
		var i6 uint = math.MaxUint64
		fmt.Println("i6:", i6)
		a1 := "hello"
		fmt.Println(reflect.TypeOf(a1))
		a2 := 3
		fmt.Println(reflect.TypeOf(a2))
		a3 := 3.0
		fmt.Println(reflect.TypeOf(a3))
		a4 := &a3
		fmt.Println("a4:", a4)
		fmt.Println(reflect.TypeOf(a4))
		var mapA map[string]string
		fmt.Println(mapA)
	*/
	/*
		const pi float64 = 3.1415926
		fmt.Println("pi:", pi)
		a := [3]int{1, 2, 3}
		var b [3]int = [3]int{}
		d := [5]int{}
		e := [...]int{1, 2, 3, 4}
		fmt.Println("a: ", a)
		fmt.Println("b: ", b)
		fmt.Println("d: ", d)
		fmt.Println("e: ", e)
		var c [3]int
		c = [3]int{4, 5, 6}
		for i := 0; i < 3; i++ {
			fmt.Println(c[i])
		}
		c = [3]int{2, 2, 2}
		for i, val := range c {
			fmt.Printf("%d\t%d\n", i, val)
		}

	*/
	a := [3][3]int{}
	a[0] = [3]int{11, 12, 13}
	a[1] = [3]int{21, 22, 23}
	a[2] = [3]int{31, 32, 33}
	fmt.Println("a:", a)
	fmt.Println("遍历一层")
	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}
	fmt.Println("遍历两层")
	for i, val := range a {
		fmt.Println("val:", val)
		for j := 0; j < len(val); j++ {
			fmt.Printf("[%d:%d]%d\t", i, j, val[j])
		}
		fmt.Println()
	}
	fmt.Println(a)
}
