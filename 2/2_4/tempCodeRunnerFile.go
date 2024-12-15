package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	num = num * num
	num = num * num
	num = num * num
	fmt.Println(num)

}