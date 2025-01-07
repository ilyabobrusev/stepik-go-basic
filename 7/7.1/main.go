package main

import "fmt"

func main() {
    name := "Саша"
    hello(name)
}

func hello(name string) {
    fmt.Println("Hello, " + name)
}

//
package main

import "fmt"

func main() {
    a := 3
    b := 6
    sum := sum(a, b) //9
    fmt.Println(sum)
}

func sum(a int, b int) int {
    sum := a + b
    return sum
}

//
package main

import "fmt"

func main() {
    hello()
}

func hello() {
    fmt.Println("Hello, PRO Go!")
}

//
package main

import "fmt"

func main() {
    var n, k int
    fmt.Scan(&n, &k)

    factN := factorial(n)
    factK := factorial(k)
    nSubtractK := n - k
    factNSubtractK := factorial(nSubtractK) // можно factorial(n - k), и тогда строчка выше не нужна

    fmt.Println(factN / (factK * factNSubtractK))
}

func factorial(number int) int {
    fact := 1
    for i := 1; i <= number; i++ {
        fact = fact * i
    }
    return fact
}

//
package main

import "fmt"

func main() {
    a := -3
    b := 2
    R := f(a)
    for t := a; t <= b; t++ {
        if f(t) > R {
            R = f(t)
        }
    }
    fmt.Println(R)
}

func f(x int) int {
    return (x + 5) * (x + 3)
}

//
package main

import "fmt"

func main() {
    i := 12
    var k int
    fmt.Scan(&k)
    for i > 0 && f(i) >= k {
        i--
    }
    fmt.Println(i)
}

func f(n int) int {
    return n*n + 20
}

//
package main

import (
	"fmt"
	
)
func main() {
	var n,k int
	fmt.Scan(&n,&k)
	fmt.Print(factorial(n)/(factorial(k)*factorial(n-k)))
}

func factorial(number int) (int){
	fact:=1

	for i:= 1 ; i <= number  ; i++ {
		fact = fact*i	
        } 
		return fact 
}

//
package main

import "fmt"

func main() {
    var a, b int
    _, _ = fmt.Scan(&a, &b)
    fmt.Println(sign(a) + sign(b))
}

func sign(x int) int {
    r := 0
    switch {
    case x > 0: r = 1
    case x < 0: r = -1
    }
    return r
}

//
package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(rank(n) * rank(m))
}

func rank(num int) int {
	cnt := 0
	for num != 0 {
		num = num / 10
		cnt++
	}
	return cnt
}
