package main

import "fmt"

func main() {
	var number int = 1024
	fmt.Println(number / 10) // 102
	fmt.Println(number % 10) // 4
}

//

package main

import "fmt"

func main(){
   var number int = 1324
   fmt.Println(number % 10) // 4
   fmt.Println(number % 100) // 24
   fmt.Println(number % 1000) // 324
   fmt.Println(number % 10000) // 1324
}

//

package main

import "fmt"

func main(){
   var number int = 1324
   fmt.Println(number / 10) // 132
   fmt.Println(number / 100) // 13
   fmt.Println(number / 1000) // 1
   fmt.Println(number / 10000) // 0
}

//

package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Println(num % 10)

}

//

package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Println(num / 10 % 10)

}

//

package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)
    fmt.Println(x % 10 + x / 10 % 10 + x / 100)
}

//

package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)
    fmt.Print(x % 10)
	fmt.Print(x / 10 % 10)
	fmt.Print(x / 100)
}

//

