package main

import (
	"fmt"
)

func main() {
	var a int
	var b string
	fmt.Scan(&a, &b)
	if a >= 12 && a <= 18 && b == "m" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

//
package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if (a > -30 && a <= -2) || (a > 7 && a <= 25) {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}

//
package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)
    if x % 2 != 0 || x % 2 == 0 && x >= 6 && x <= 20 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

//
package main
import "fmt"

func main(){
    var  num float64
    _,_ = fmt.Scan(&num)
    
    if  num != 0 {
        fmt.Print(1/num)
    } else {
        fmt.Print("Обратного числа не существует")
    }
}

//
package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    switch {
        case n <= 3: 
		    fmt.Print("начинающий")
        case n > 3 && n <= 7: 
		    fmt.Print("младший разработчик")
        case n > 7 && n < 16: 
		    fmt.Print("средний разработчик")
        case n > 16: 
		    fmt.Print("старший разработчик")
    }
}

//
package main
import "fmt"
func main() {
    var d1, d2, d3 int
    fmt.Scan(&d1, &d2, &d3)
    if d3 > d1 + d2 {
        fmt.Println(d1 * 2 + d2 * 2)
    } else if d1 > d2 + d3 {
        fmt.Println(d2 * 2 + d3 * 2)
    } else if d2 > d1 + d3 {
        fmt.Println(d1 * 2 + d3 * 2)
    } else {
        fmt.Println(d1 + d3 + d2)
    }
}
