package main

import "fmt"

func main(){
	var x, y int
	fmt.Scan(&x, &y)

	if x > 5 && y < 10 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

//
package main

import "fmt"

func main(){
	var x, y int
	fmt.Scan(&x, &y)

	if x > 5 || y < 10 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

//
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
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
	var x int
	fmt.Scan(&x)
	if 100 <= x && x <= 999 {
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
	var x int
	fmt.Scan(&x)
	if x > -1 && x < 17 {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}

//
package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	if x <= -3 || x >= 7 {
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
    c := x % 10
    b := (x / 10) % 10
    a := (x / 100) % 10
    if c != b && c!= a && b!= a {
        fmt.Println("YES")
    }else { 
        fmt.Println("NO")
    }
}

//
package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	a := x % 1000
	b := x / 1000

	if (a%10)+(a/100)+(a/10%10) != (b%10)+(b/100)+(b/10%10) {
		fmt.Println("NO")

	} else {
		fmt.Println("YES")
	}

}

//
package main

import "fmt"

func main() {
    var n int16
    _,_ = fmt.Scan(&n)
    
    if (n % 4 == 0 && n % 100 != 0) || n % 400 == 0{
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
	var x int
	fmt.Scan(&x)
	if x >= 0 && x <= 13 {
		fmt.Println("детство")
	} else if x >= 14 && x <= 24 {
		fmt.Println("молодость")
	} else if x >= 25 && x <= 59 {
        fmt.Print("зрелость")
    } else {
        fmt.Print("старость")
    }
}

//

package main

import "fmt"

func main() {
    var answer = "NO"
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)

    if a == c || b == d {
        answer = "YES"
    }

    fmt.Println(answer)
}

//

package main

import "fmt"

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)
    if (a - b == c - d) || (a + b == c + d) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

//

