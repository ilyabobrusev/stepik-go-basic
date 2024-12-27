package main

import (
	"fmt"
)

func main() {
	n := 21
	i := 1
	cnt := 0
	for i <= n {
		if i%10 == 0 {
			cnt = 3
		} else {
			cnt = cnt + 5
		}
		i++
	}
	fmt.Println(cnt)
}

//
package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)

    i := 2
    for n%i != 0 {
        i++

    }
    fmt.Println(n / i)
}

//

package main

import "fmt"

func main() {
    var a int
    _,_ = fmt.Scan(&a)
    count := 0    
    for a % 3 == 0 {
        a /= 3
        count++
    }
    fmt.Print(count)
}

//
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	value := 1
	for value <= n {
		fmt.Println(value)
		value = value * 2
	}
}

//
package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    k := 0
    for i := 1; i < n; i *= 2 {
        k++
    }

    fmt.Println(k)
}

//
package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    for a % b != 0 {
        a, b = b, a % b
    }
    fmt.Println(b)
}

//
