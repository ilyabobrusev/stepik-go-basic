package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0    // сумматор
	for n > 0 { // пока цифры числа не закончатся
		digit := n % 10   // последняя цифра числа
		sum = sum + digit // складываем в сумматор
		n = n / 10        // избавляемся от последней цифры.
	}

	fmt.Println(sum)
}

//
package main

import (
    "fmt"
)

func main() {
    n := 6808
    cnt := 1
    for n > 0 {
        cnt = cnt + 1
        n = n / 10
    }
    fmt.Println(cnt)
}

//
package main
import "fmt"
func main() {
    var n int; fmt.Scan(&n)
    cnt:=0
    for n>0 {
        if n%10==4 {cnt++}
        n/=10
    }
    fmt.Println(cnt)
}

//
package main

import (
	"fmt"
)

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	sum := 0
	for i := a; i != 0; {
		sum += i % 10
		i /= 10
	}
	if a%sum == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

//
package main 
import "fmt"
func main () {
    var n int
    fmt.Scan(&n)
    for n>0 {
        fmt.Print(n%10)
        n/=10
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

    for n > 0 {
        fmt.Println(n % 2)
        n = n / 2
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

    for n > 0 {
        fmt.Print(n % 2)
        n = n / 2
    }
}

//
package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    i := 0
    for n != 0 {
        i = i + n%2
        n /= 2
    }
    fmt.Println(i)
}
