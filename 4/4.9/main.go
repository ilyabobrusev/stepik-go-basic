package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	for x := 1; x <= 10000; x++ {
		if a*x*x+b*x+c == 0 {
			fmt.Println(x)
		}
	}
}

//
package main

import (
    "fmt"
)

func main() {
    for i := 1; i < 10; i++ {
        for j := 0; j < 10; j++ {
            if i != j {
                fmt.Println(i*10 + j)
            }
        }
    }
}

//
package main

import (
    "fmt"
)

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)
    
    for x:=0; x<=30000; x++ {
        if a*x*x*x + b*x*x + x*c + d == 0 {
            fmt.Print(x)
            fmt.Print(" ")
        }
    }
}

//
package main

import "fmt"

func main() {
	var n, c, d int
	_, _ = fmt.Scan(&n, &c, &d)
	if c%d != 0 {
		for i := c; i <= n; i += c {
			if i%d != 0 {
				fmt.Println(i)
			}
		}
	}
}

//
package main

import "fmt"

func main() {
	count := 0
	for i := 100; 1000 > i; i++ {
		if i%10+i/100+i/10%10 == 8 {
			count++
		}
	}
	fmt.Print(count)

}

//
package main

import "fmt"

func main() {
    for i := 10; i <= 100; i++ {
       if i == 2*(i/10)*(i%10) {
          fmt.Print(i, " ")
       }
    }
}

//
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10000; i++ {
		sum := 0
		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if i == sum {
			fmt.Printf("%d,", i)
		}
	}
}

//
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	sum := 0
	for x := 0; x <= 1000; x++ {
		if (x-e) != 0 && ((a*x*x*x)+(b*x*x)+(c*x)+d) == 0 {
			sum++
		}

	}
	fmt.Println(sum)
}
