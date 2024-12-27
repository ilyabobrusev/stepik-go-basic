package main

import "fmt"

func main() {
    sum := 0 // переменная - сумматор. В ней будет храниться сумма.
    for i := 1; i <= 100; i++ {
        sum += i // прибавляем к текущей сумме очередной элемент и результат снова записываем в сумматор
    }
    fmt.Println(sum)
}
//

package main

import "fmt"

func main() {
    sum := 0
    var number int

    for i := 1; i <= 10; i++ {
        fmt.Scan(&number) //вводим число, каждый раз перезаписывая значение переменной
        sum += number     // суммируем
    }

    fmt.Println(sum)
}
//

package main

import "fmt"

func main() {
	sum := 0
	for i := 15; i > 1; i-- {
		sum = sum + i
	}
	fmt.Println(sum)
}

//
package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    count := 0
    for i := 1; i <= n; i++ {
        if i%3 == 0 && i%7 != 0 {
            count++
        }
    }

    fmt.Println(count)
}

//
package main

import "fmt"

func main() {
    var n, count int
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        if n % i == 0 {
            count++
        }
    }
    fmt.Println(count)
}

//

package main

import (
	"fmt"
)

func main() {
	var n, a int
	fmt.Scan(&n)
	count := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		count += a
	}
	fmt.Println(count)
}

//
package main

import ("fmt")

func main() {
    var a,b int
    fmt.Scan(&a)
    sum:=0
    for i:=1; i<=a; i++ {
        fmt.Scan(&b)
        if b%2 ==0 && b%3 != 0 {
        sum+=b
        }
    }
    fmt.Print(sum)
}

//
package main

import "fmt"

func main() {
    var n, number int
    fmt.Scan(&n)
        
    for i := 0; i < n; i++ {
        var x int
        if fmt.Scan(&x); x % 10 == 0 {
        number++
        }
    }
    fmt.Println(number)
}

//
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	count := 0

	for i := 0; i <= a; i++ {
		fmt.Scan(&b)
		if b == 0 {
			count++
		}
	}
	if count > 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}

