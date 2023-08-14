package main

import (
	"fmt"
)

var totalOperations int
var totalDigits int

func calculator() float64 {
	var number1, number2 float64
	var result float64
	var operator string

	totalOperations++
	totalDigits++

	fmt.Println("計算機已啟動")

	fmt.Print("請插入第一個數字 : ")
	fmt.Scan(&number1)

	fmt.Print("請插入第二個數字 : ")
	fmt.Scan(&number2)

	fmt.Print("請插入數學運算 ( + - * / ) : ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result = number1 + number2
		fmt.Printf("%.0f + %.0f = %.0f\n", number1, number2, result)
	case "-":
		result = number1 - number2
		fmt.Printf("%.0f - %.0f = %.0f\n", number1, number2, result)
	case "*":
		result = number1 * number2
		fmt.Printf("%.0f * %.0f = %.0f\n", number1, number2, result)
	case "/":
		if number2 == 0.0 {
			fmt.Println("除以 0 是未定義的")
		} else {
			result = number1 / number2
			fmt.Printf("%.0f / %.0f = %.0f\n", number1, number2, result)
		}
	default:
		fmt.Println("表達錯誤")
	}

	fmt.Printf("結果： %f\n", result)
	return result
}

func main() {
	for {
		var choice string
		fmt.Println("開始新的循環嗎? \n y/n")
		fmt.Scan(&choice)

		if choice == "y" {
			calculator()
			fmt.Printf("繼續\n")
		} else if choice == "n" {
			fmt.Printf("\n謝謝您的合作. 計算機關閉")
			fmt.Printf("\n一共操作: %d\n", totalOperations)
			fmt.Printf("一共輸入的數字: %d\n", totalDigits)
			break
		} else {
			fmt.Println("表達錯誤")
		}
	}
}
