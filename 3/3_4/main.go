package main

import (
	"fmt"
)

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
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
	
	if a%2 != 0 {
		a++
		fmt.Println(a)
	} else {
		a++
		a++
		fmt.Println(a)
	}
}

//
package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)

	var digit0 = x % 10
	var digit1 = x / 10 % 10
	var digit2 = x / 100 % 10
	var digit3 = x / 1000

	if digit0 == digit3 && digit1 == digit2 {
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
	
	b3 := a%10
	b2 := a/10%10
	b1 := a/100

	if b1 < b2 && b2 < b3 {
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
	var a, b int
	fmt.Scan(&a, &b)

	if a%b == 0 {
		c := a/b
		fmt.Println(c)
	} else {
		c := a/b+1
		fmt.Println(c)
	}
}

//
package main

import (
	"fmt"
)

func main() {
	var k2, k3, k5, k6 int
	fmt.Scan(&k2, &k3, &k5, &k6)
	n256 := min(k2, k5, k6)
	n32 := min(k2-n256, k3)
	fmt.Println(n256*256 + n32*32)
}

//
package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)
	if x1+2 == x2 || x1-2 == x2 {
		if y1+1 == y2 || y1-1 == y2 {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	} else if x1+1 == x2 || x1-1 == x2 {
		if y1+2 == y2 || y1-2 == y2 {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	} else {
		fmt.Print("NO")
	}
}

//
package main

import ("fmt")

func main() {
    var h1, w1, h2, w2 int
    fmt.Scan(&h1, &w1, &h2, &w2)
    dh := h1 - h2
    dw := w1 - w2
    
    if dh*dh == dw*dw && dh != 0 || dh == 0 || dw == 0  {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

//
package main

import "fmt"

func main() {
    var a, b, c float64
    fmt.Scan(&a, &b, &c)
    
    if a + b > c && c + a > b && b + c > a {
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
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	a = (a + 1) / 2
	b = (b + 1) / 2
	c = (c + 1) / 2
	
	fmt.Print(a + b + c)
}

//
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a > 2 && a%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
