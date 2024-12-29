package main

import (
	"fmt"
)

func main() {
	var n, x, count int
	count = 0
	for fmt.Scan(&n); n > 0; n-- {
		fmt.Scan(&x)
		if x == 0 {
			count++
		}
	}
	fmt.Println(count)
}

//
package main
import "fmt"

func main(){
    var num, kopy, rev int
    res := "NO"
    fmt.Scan(&num)
    kopy = num
    
    for num > 0 {
        rev = rev*10 + num % 10
        num /= 10
    }
    if kopy == rev {
        res = "YES"
    }
    fmt.Print(res)
}

//
package main

import "fmt"

func main(){
    var n int
    fmt.Scan(&n)
    i := 1
    res := 0
    for n > 0 {
        if n % 10 != 5 && n % 10 != 7 {
            res += (n % 10) * i
            i = i * 10
        }
        n /= 10
    }
    fmt.Println(res)
}

//
package main

import (
	"fmt"
	"math"
)

func main(){
    var a, n float64
    fmt.Scan(&a, &n)
    fmt.Println(math.Pow(a, n))
}

//
package main

import "fmt"

func main() {
    var number, n int
    fmt.Scan(&n, &number)
	min := number
    cnt := 1
    for i := 1; i < n; i++ {
        fmt.Scan(&number)
        if number < min {
            min = number
            cnt = 1
        } else if min == number {
            cnt += 1
        }
    }
    fmt.Println(cnt)
}

//
package main

import (
	"fmt"
	"math"
)

func main(){
    var a float64
    fmt.Scan(&a)
    fmt.Println(math.Sqrt(a))
}

//
package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	for number > 9 {
		sum := 0
	    for number > 0 {
		    sum += number % 10
		    number /= 10
	    }
        
        number = sum
	}

	fmt.Println(number)
}

//
package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("NO")
}

//
package main

import "fmt"

func main() {
   var year int
    fmt.Scan(&year)
    niceYear := year + 1
    for i := niceYear; i > 0; i /= 10 {
        for j := i / 10; j > 0; j /= 10 {
            if i % 10 == j % 10 {
                niceYear++
                i = niceYear * 10
                break
            } 
        }
    }
    fmt.Print(niceYear)
}
