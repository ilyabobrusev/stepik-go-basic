package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for n != 0 { // пока число, хранимое в переменной n не равно нулю
		fmt.Scan(&n)
	}
}

//
package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    sum := 0

    for n != 0 {
        sum = sum + n // сначала прибавляем, а потом вводим!
        fmt.Scan(&n)
    }

    fmt.Println(sum)
}

//
