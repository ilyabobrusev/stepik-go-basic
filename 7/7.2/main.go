package main

import "fmt"

func sum(n int) int {
    s := 0
    for n > 0 {
        s += n % 10
        n /= 10
    }
    return s
}

func f(n int) bool{
    const c int = 1000
    return sum(n / c) == sum(n % c)
}

func main() {
    var a, b, c int
    fmt.Scan(&a, &b)
    
    if f(a) && f(b) {
        c = 1
    } else {
        c = -1
    }
    
    fmt.Println(c)
}

//
package main

import (
    "fmt"
    
)

func avg(n int) float64 {
    return float64(1 + n) / 2
}


func main() {
    var n, m int
    fmt.Scan(&n, &m)
    fmt.Println(avg(n) + avg(m))
}


//
package main

import "fmt"

func main() {
    var a,b int
    _,_ = fmt.Scan(&a, &b)
    
    fmt.Print(obr(a)+obr(b))
}


func obr(num int) int {
    res:=0
    for num>0 {
        res *=10
        res +=num%10
        num /=10
    }
    return res
}

//
package main

import (
    "fmt"    
)

func Fact2(number int) int {
	total := 1   
    for i := number; i > 1; i -= 2 {
		total *= i
	}
	return total
}

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    fmt.Println(Fact2(a), Fact2(b), Fact2(c))
}

//
package main

import "fmt"

func f() int {
    var a , b int
    fmt.Scan(&a, &b)
    
    c := b
    
    for i := 1 ; i < a ; i++  {
        fmt.Scan(&b)
        if b > c { c = b }
    }
    
    return c
}

func main(){
    fmt.Println( f() * f() )
}

//
package main

import "fmt"

func main() {
    var a,b string
    fmt.Scan(&a,&b)
    fmt.Println(compareLen(a,b))
}

func compareLen(x,y string) int {
    switch {
        case len(x) == len(y):
        return 0
        case len(x) > len(y):
        return 1
        default:
        return 2
    }
}

//
package main

import "fmt"

func main() {
    var a,b string
    fmt.Scan(&a,&b)
    fmt.Println(compareLen(a,b))
}

func compareLen(x,y string) int {
    switch {
        case len(x) == len(y):
        return 0
        case len(x) > len(y):
        return 1
        default:
        return 2
    }
}

//
package main
import "fmt"

func main(){
    var n, m int
    fmt.Scan(&n, &m)
    
    n1 := sum(n)
    m1 := sum(m)
    
    switch{
    case n1 > m1:
        fmt.Print(1)
    case n1 < m1:
        fmt.Print(2)
    case n1 == m1:
        fmt.Print(0)
    }
}

func sum(s int)int{
    s1 := 0
    for s > 0{
        s1 = s%10 + s1
        s /= 10
    }
    return s1
}


//
package main

import "fmt"

func main() {
    var s1, s2 string
    _,_ = fmt.Scan(&s1,&s2)
    fmt.Print(c_b(s1) + c_b(s2))
}

func c_b(s string) int {
    cnt:=0
    for i:=0; i<=len(s)-1; i++ {
        if string(s[i]) == "b"{
            cnt++
        }
    }
    return cnt
}

//
package main

import (
    "fmt"
)

func main() {
    var n int
    _, _ = fmt.Scan(&n)
    fmt.Println(primeN(n))
}

func primeN(n int) string {
    for i := 2; i < n; i++ {
        if n % i == 0 {
            return "composite"
        }
    }
    return "prime"
}

//
package main

import (
	"fmt"
)

func main() {
	var n, m, o int
	fmt.Scan(&n, &m, &o)
	fmt.Println(sumrange(n, m) + sumrange(m, o))
}

func sumrange(x, y int) int {
	sum := 0
	for i := x; i <= y; i++ {
		sum = sum + i
	}
	return sum
}

//
