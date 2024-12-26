package main

import "fmt"

func main() {
	for i := 0; i < 19; i++ {
		fmt.Println("Hello, PRO Go!")
	}
}

//

package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
}

//
package main

import "fmt"

func main() {
    for i := 2; i <= 100; i = i + 2 {
        fmt.Println(i)
    }
}

//

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

    for i := a; i <= b; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
}

//
package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}
}

//
package main

import "fmt"

func main() {
    var a, b int 
    fmt.Scan(&a, &b)
    for ; b >= a; b-- {
        fmt.Println(b)
    } 
}

//
package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		fmt.Println(i*i)
	}
}