package main

import "fmt"

func main() {
	var (
		str string
		k   int
	)
	fmt.Scan(&str, &k)
	if k > len(str) {
		fmt.Println("NO")
		return
	}
	fmt.Println(string(str[k-1]))
}

//
package main

import (
    "fmt"
)

func main() {
    var s1, s2 string

    _, _ = fmt.Scan(&s1, &s2)
    
    c1 := []rune(s1)[0]
    c2 := []rune(s2)[0]
    if c1 > c2 {
        c1, c2 = c2, c1
    }

    for i := c1; i <= c2; i++ {
        fmt.Printf("%c ", i)
    }
}

//
package main
import "fmt"

func main() {
    var a,b,c string
    fmt.Scan(&a,&b,&c)
    fmt.Print(a, " ", string(b[0]), ".", string(c[0]), ".")
}

//
package main

import "fmt"

func main() {
    var a, b string
    fmt.Scan(&a)
    for _, err := fmt.Scan(&b); err == nil; _, err = fmt.Scan(&b) {}
    if a[0] == b[len(b) - 1] {
        fmt.Print("YES")   
    } else {
        fmt.Print("NO")   
    }
    
}

//
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	c := []rune(s)
	cnt := 0
	for i:=0; i < len(c); i++ {
		switch c[i] {
		case 'a', 'e', 'i', 'o', 'u':
			cnt++
		default:
			continue
		}
	}
	fmt.Println(cnt)
}

//
package main
import (
    "fmt"
)

func main() {
    var b string 
    fmt.Scan(&b)
    a := []rune(b)
    for i := 0; i < len(a); i++ {
        if string(a[i]) == "e" {
            fmt.Print("i")
            continue
        }
        fmt.Print(string(a[i]))
    }
}

//// v 2
package main

import (
	"fmt"
	"strings"
)

func main() {
	var string1 string
	_, _ = fmt.Scan(&string1)
	fmt.Print(strings.ReplaceAll(string1, "e", "i"))
}

//
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
    str := scanner.Text()
    
	for i := 0; i < len(str); i++ {
		if str[i] >= 48 && str[i] <= 57 {
			fmt.Print(string(str[i]), " ")
		}
	} 
}

//
package main
import("fmt")
func main(){
    var(
        s string
        n int)
    fmt.Scan(&n)
    for i:=0; i<n; i++{
        fmt.Scan(&s)
        if len(s)>10{
            fmt.Print(string(s[0]), len(s)-2, string(s[len(s)-1]))
            fmt.Println()
        }else {
            fmt.Println(s)
        }
    }
}

//
package main

import (
	"fmt"
)

func main() {
	 var n int 
	 var str string
	 fmt.Scan(&n, &str)
     kolX := 0
	 for i := 1; i < n-1; i++ {
		if str[i] == 'x' && str[i-1] == 'x' && str[i+1] == 'x' {
           kolX++
		}
	 }
	 fmt.Print(kolX)
}

//
package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "." {
			fmt.Print("0")
		} else if string(str[i]) == "-" && string(str[i+1]) == "." {
			fmt.Print("1")
			i++
		} else if string(str[i]) == "-" && string(str[i+1]) == "-" {
			fmt.Print("2")
			i++
		}
	}
}

//// v 2
package main

import (
	"fmt"
    "strings"
)

func main() {
    var str string
    fmt.Scan(&str)
    str=strings.ReplaceAll(str, "--", "2")
    str=strings.ReplaceAll(str, "-.", "1")
    str=strings.ReplaceAll(str, ".", "0")
    fmt.Print(str)
}

//
