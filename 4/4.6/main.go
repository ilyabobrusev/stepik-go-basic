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
package main 
import "fmt"
func main () {
    var n,sum int
    for fmt.Scan(&n) ; n!=0; fmt.Scan(&n) {
        if n%2==0 && n%3!=0 {
            sum+=n
        }
    }
    fmt.Print(sum)
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
        if n>0{
        sum = sum + 1
        } 
        if n<0 {
        sum = sum - 1
        }
        fmt.Scan(&n)
    }

    fmt.Println(sum)
}

//
package main
import "fmt"
func main () {
    var n, dif, i float64
    for fmt.Scan(&n); n!=0; fmt.Scan(&n) {
        dif+=n
        i++
    }
    fmt.Print(dif/i)
}

//
package main 
import ("fmt")
func main() {
    var n, i, sum int
    fmt.Scan (&n)
    fmt.Scan (&i)
    for i != 0{
        if i > n {
            sum++}
            n, i  = i, n
            fmt.Scan(&i)
        
    }
    fmt.Println(sum)
}

//
package main

import "fmt"

func main() {
	var prev, next int
	_, _ = fmt.Scan(&prev, &next)
	var count = 0

	for next != 0 {
		if next*prev < 0 {
			count++
		}

		prev = next
		_, _ = fmt.Scan(&next)
	}

	fmt.Println(count)
}

//
