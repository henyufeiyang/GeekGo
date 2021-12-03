package main

import "fmt"

func main() {
	var i int = 0
	var sum float64
	for i < 3 {
		var name string
		fmt.Print("姓名：")
		fmt.Scanln(&name)

		var weight float64
		fmt.Print("体重(千克)：")
		fmt.Scanln(&weight)

		var high float64
		fmt.Print("身高(米)：")
		fmt.Scanln(&high)

		var age int
		fmt.Print("年龄(岁)：")
		fmt.Scanln(&age)

		var sex string
		fmt.Print("性别(man/woman)：")
		fmt.Scanln(&sex)

		var sexwight int
		if sex == "man" {
			sexwight = 1
		} else {
			sexwight = 0
		}

		var bmi, bfr float64
		bmi = weight / (high * high)
		bfr = 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexwight)
		// fmt.Printf("bmi: %f,bfr: %f", bmi, bfr)

		if sex == "man" {
			switch {
			case age >= 18 && age <= 39:
				if bfr <= 10 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要增加体重了\n", name, bmi, bfr)
				} else if bfr > 10 && bfr <= 16 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你的身材很标准\n", name, bmi, bfr)
				} else if bfr > 16 && bfr <= 21 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你超重了，需要控制饮食\n", name, bmi, bfr)
				} else if bfr > 21 && bfr <= 26 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要安排健身了\n", name, bmi, bfr)
				} else {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要马上行动去减肥吧\n", name, bmi, bfr)
				}
			case age >= 40 && age <= 59:
				if bfr <= 11 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要增加体重了\n", name, bmi, bfr)
				} else if bfr > 11 && bfr <= 17 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你的身材很标准\n", name, bmi, bfr)
				} else if bfr > 17 && bfr <= 22 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你超重了，需要控制饮食\n", name, bmi, bfr)
				} else if bfr > 22 && bfr <= 27 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要安排健身了\n", name, bmi, bfr)
				} else {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要马上行动去减肥吧\n", name, bmi, bfr)
				}
			case age >= 60:
				if bfr <= 13 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要增加体重了\n", name, bmi, bfr)
				} else if bfr > 13 && bfr <= 19 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你的身材很标准\n", name, bmi, bfr)
				} else if bfr > 19 && bfr <= 24 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你超重了，需要控制饮食\n", name, bmi, bfr)
				} else if bfr > 24 && bfr <= 29 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要安排健身了\n", name, bmi, bfr)
				} else {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要马上行动去减肥吧\n", name, bmi, bfr)
				}
			default:
				fmt.Println("%s 你的BMI是 %f，BFR是 %f，你还在长身体\n", name, bmi, bfr)
			}
		} else {
			switch {
			case age >= 18 && age <= 39:
				if bfr <= 20 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要增加体重了\n", name, bmi, bfr)
				} else if bfr > 20 && bfr <= 27 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你的身材很标准\n", name, bmi, bfr)
				} else if bfr > 27 && bfr <= 34 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你超重了，需要控制饮食\n", name, bmi, bfr)
				} else if bfr > 34 && bfr <= 39 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要安排健身了\n", name, bmi, bfr)
				} else {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要马上行动去减肥吧\n", name, bmi, bfr)
				}
			case age >= 40 && age <= 59:
				if bfr <= 21 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要增加体重了\n", name, bmi, bfr)
				} else if bfr > 21 && bfr <= 28 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你的身材很标准\n", name, bmi, bfr)
				} else if bfr > 28 && bfr <= 35 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你超重了，需要控制饮食\n", name, bmi, bfr)
				} else if bfr > 35 && bfr <= 40 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要安排健身了\n", name, bmi, bfr)
				} else {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要马上行动去减肥吧\n", name, bmi, bfr)
				}
			case age >= 60:
				if bfr <= 22 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要增加体重了\n", name, bmi, bfr)
				} else if bfr > 22 && bfr <= 29 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你的身材很标准\n", name, bmi, bfr)
				} else if bfr > 29 && bfr <= 36 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你超重了，需要控制饮食\n", name, bmi, bfr)
				} else if bfr > 36 && bfr <= 41 {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要安排健身了\n", name, bmi, bfr)
				} else {
					fmt.Printf("%s 你的BMI是 %f，BFR是 %f，你需要马上行动去减肥吧\n", name, bmi, bfr)
				}
			default:
				fmt.Println("%s 你的BMI是 %f，BFR是 %f，你还在长身体\n", name, bmi, bfr)
			}
		}
		sum += bfr
		i++
	}
	var bfravg float64
	bfravg = sum / float64(i)
	fmt.Printf("总人数： %d，平均体脂率：%f", i, bfravg)
}
