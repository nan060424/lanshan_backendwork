package main

import "fmt"

func add(a, b int) int { return a + b }
func main() {
	var num1, num2 int
	fmt.Println("请输入一个数")
	fmt.Scan(&num1)
	fmt.Println("请输入一个数")
	fmt.Scan(&num2)
	result := add(num1, num2)
	fmt.Println(result)
}
