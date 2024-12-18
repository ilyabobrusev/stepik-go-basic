package main

import "fmt"

func main() {
	a := 5.0 // наличие точки дает понять компилятору что это вещественное число
	b := 2.0

	sum := a + b
	diff := a - b
	mult := a * b

	fmt.Println(a, " + ", b, " = ", sum)
	fmt.Println(a, " - ", b, " = ", diff)
	fmt.Println(a, " * ", b, " = ", mult)
}

//

package main

import "fmt"

func main(){
   var n int = 5 
   var floatDivider float32 = 2.0
   fmt.Println(float32(n) / floatDivider) // 2.5
}

//

package main

import "fmt"

func main() {
   f := 0.1
   f += 0.2
   fmt.Println(f) // Выводит: 0.30000000000000004
}

//

package main

import "fmt"

func main() {
   f := 5.8
   fmt.Println(f - 5.0) // Выводит: 0.7999999999999998
}

//

package main

import "fmt"

func main(){
   var d float32 = 2.9
   fmt.Println(int(d)) // 2
}

//

package main

import "fmt"

func main() {
    var r float32
    fmt.Scan(&r)
    fmt.Print(r * r * 3.14)
}

//

package main

import "fmt"

func main() {
    var a, b float64
    fmt.Scan(&a, &b)
    fmt.Print(0.5 * a * b)
}

//

package main

import "fmt"

func main() {
    var f float64
    fmt.Scan(&f)
    fmt.Print((f - 32) * 5 / 9)
}

//

package main

import "fmt"

func main() {
    var a, b float64
    fmt.Scan(&a, &b)
    fmt.Print((a + b) / 2)
}

//

package main

import "fmt"

func main() {
    var n float64
    fmt.Scan(&n)
    fmt.Println(n - float64(int(n)))
}

//

package main

import (
    "fmt"
    "math"
)

func main(){
    var n float64 = -9
    fmt.Println(math.Abs(n)) // берет модуль числа и выводит на экран число 9
}

//

package main

import (
    "fmt"
    "math"
)

func main() {
    var n float64 = 8
    fmt.Println(math.Sqrt(n))
}

//

package main

import (
    "fmt"
    "math"
)

func main() {
    var n float64 = 8
    fmt.Println(math.Pow(n, 2))
}

//

package main

import (
    "fmt"
    "math"
)

func main() {
    var n float64 = 8
    fmt.Println(math.Max(n, 4)) // выводит на экран число 8
    fmt.Println(math.Min(n, 4)) // выводит на экран число 4
}

//

package main

import (
    "fmt"
    "math"
)

func main() {
    var n float64
    fmt.Scan(&n)
    fmt.Println(n / math.Pow(2, 13))
}

//


package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b float64
    _, _ = fmt.Scan(&a, &b)
    fmt.Println(math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2)) + (a + b))
}

//

package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b, c, d float64
    _, _ = fmt.Scan(&a, &b, &c, &d)
    fmt.Println(math.Sqrt(math.Pow(a - c, 2) + math.Pow(b - d, 2)))
}

//

package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b, c, d float64
    _, _ = fmt.Scan(&a, &b, &c, &d)
    fmt.Println(math.Abs((a-c)) + math.Abs((b-d)))
}

//

package main

import (
    "fmt"
)

func main() {
    var a, b, c int
    _, _ = fmt.Scan(&a, &b, &c)
    fmt.Println((a+b+c), (a*b*c))
}

//

package main

import (
    "fmt"
)

func main() {
    var n int
    _, _ = fmt.Scan(&n)
    fmt.Println(n / 100 % 10)
}

//

package main

import (
    "fmt"
)
func main() {
    var a int
    _, _ = fmt.Scan(&a)
    fmt.Println((a%10)+(a/10%10)+(a/100%10))
}

//

package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b float64
    _, _ = fmt.Scan(&a, &b)
    fmt.Println(math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2)))
}
