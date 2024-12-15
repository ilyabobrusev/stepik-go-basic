package main

import "fmt"

func main() {
	var a int = 6
	var b int = a * 8 // 48
	a = b - 8         // 40. Заметьте, что тип данных мы не написали
	var c int = b + a // 88
	fmt.Println(c)    // вывод 88
}
