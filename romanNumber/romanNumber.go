package main

import "fmt"

func romanNumber(input int) string {

	result := ""
	romanIntValue := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanString := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for index := 0; index < len(romanIntValue); index++ {
		for romanIntValue[index] <= input {
			input = input - romanIntValue[index]
			result = result + romanString[index]
		}
	}

	return result
}

func main() {

	result := romanNumber(1)
	fmt.Printf("%#v\n", result)
	result = romanNumber(4)
	fmt.Printf("%#v\n", result)
	result = romanNumber(5)
	fmt.Printf("%#v\n", result)
	result = romanNumber(9)
	fmt.Printf("%#v\n", result)
	result = romanNumber(10)
	fmt.Printf("%#v\n", result)
	result = romanNumber(40)
	fmt.Printf("%#v\n", result)
	result = romanNumber(50)
	fmt.Printf("%#v\n", result)
	result = romanNumber(90)
	fmt.Printf("%#v\n", result)
	result = romanNumber(100)
	fmt.Printf("%#v\n", result)
}
