// package main

// import "fmt"

// func main() {
// 	var num int
// 	_, _ = fmt.Scan(&num)
// 	fmt.Println(num * num)
// }

// #

// package main

// import "fmt"

// func main() {
// 	var num int
// 	_, _ = fmt.Scan(&num)
// 	fmt.Println(num / 1000)
// }

// #

// package main

// import "fmt"

// func main() {
// 	var num int
// 	_, _ = fmt.Scan(&num)
// 	num = num * num * num
// 	fmt.Println(num * num)

// }

// #

// package main

// import "fmt"

// func main() {
// 	var num1, num2, num3 int
// 	_, _ = fmt.Scan(&num1, &num2, &num3)
// 	fmt.Println(num1 * num2 * num3)

// }

// #

// package main

// import "fmt"

// func main() {
// 	var n, k int
// 	_, _ = fmt.Scan(&n, &k)
// 	fmt.Println(k / n)

// }

// #

// package main

// import "fmt"

// func main() {
// 	var n, k int
// 	_, _ = fmt.Scan(&n, &k)
// 	fmt.Println(k % n)

// }

// #

// package main

// import "fmt"

// func main() {
// 	var num int
// 	_, _ = fmt.Scan(&num)
// 	fmt.Println(num)
// 	num++
// 	fmt.Println(num)
// 	num++
// 	fmt.Println(num)
// }

// #

package main

import "fmt"

func main() {
	var t, c, p, h int
	_, _ = fmt.Scan(&t, &c, &p, &h)
	fmt.Println(3 * (t + c + p + h))
}
