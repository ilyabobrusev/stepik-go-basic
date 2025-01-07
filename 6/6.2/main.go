package main

import (
    "fmt"
)

func main() {
    emoji := []rune("Доброго дня 😀")

    for i := 0; i < len(emoji); i++ {
        fmt.Println(emoji[i], string(emoji[i])) // выводим код символа и его строковое представление
    }
}

//
package main

import (
    "fmt"
)

func main() {
    symb1 := 'a' // 97
    symb2 := 'b' // 98
    if symb1 > symb2 {
        fmt.Println("a > b")
    } else {
        fmt.Println("a < b") // Выведется в консоль
    }
}

//
package main

import (
    "fmt"
)

func main() {
    var c rune = 'B'

    if c >= 'A' && c <= 'Z' { //если символ лежит в диапазоне от 'A' до 'Z' включительно
        fmt.Println("Заглавная буква английского алфавита")
    }
}

//
package main

import (
    "fmt"
)

func main() {
    var c rune = 'a'
    fmt.Println(string(c + 3)) // 'd'

    var b rune = '0'
    fmt.Println(string(c + b))
    fmt.Println(string(c - b))
    fmt.Println(string(c * b))
    fmt.Println(string(c / b))
    fmt.Println(string(c % b))
}

//
package main

import (
    "fmt"
)

func main() {
    for w := 'A'; w <= 'Z'; w++ {
        fmt.Println(string(w))
    }
}

//
package main

import (
	"fmt"
)

func main() {
	var c string
	fmt.Scan(&c)
	rs := []rune(c)
	if rs[0] >= '0' && rs[0] <= '9' {
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
	var c string
	fmt.Scan(&c)
	rs := []rune(c)
    for w := 'a'; w <= rs[0]; w++ {
        fmt.Print(string(w), " ")
    }
}

//
package main

import (
    "fmt"
)

func main() {
	var c string
	fmt.Scan(&c)
	rs := []rune(c)
    for w := rs[0]; w <= 'z'; w++ {
        fmt.Print(string(w), " ")
    }
}

/////// v2
package main

import (
    "fmt"
)

func main() {
	var c string
	fmt.Scan(&c)
    for w := c[0]; w <= 'z'; w++ {
        fmt.Print(string(w), " ")
    }
}

//
package main

import (
	"fmt"
)

func main() {
	
   var c rune 
   fmt.Scanf("%c", &c)

   if c >= 'a' && c <= 'z' {
      fmt.Println(string(c - 32))
   } else {
      fmt.Println(string(c + 32))
   }
}