package main

import (
    "fmt"
)

func main() {
    var s string
	fmt.Print("Введите имя: ")
    fmt.Scan(&s)

    fmt.Println(s)
}

//
package main

import (
    "fmt"
)

func main() {
    s := "PROGo"
    countChar := len(s)

    for index := 0; index < countChar; index++ {
        fmt.Println(string(s[index]))
    }
}

//
package main

import (
    "fmt"
)

func main() {
    var s string
    fmt.Scan(&s)
    countChar := len(s)

    countA := 0
    for index := 0; index < countChar; index++ {
        if s[index] == 'A' {
            countA++
        }
    }

    fmt.Println(countA)
}

//
package main

import ( "fmt"
        "strings"
        "bufio"
        "os"
       )

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    _ = scanner.Scan()
    text := scanner.Text() 
    words := strings.Split(text, " ")
    
    fmt.Print(len(words))
    
}

//
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str, res string
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	str = sc.Text()
	for _, a := range str {
		res = string(a) + res
	}
	fmt.Println(res)
}

//
package main

import "fmt"

func main(){
    var (
        str string
        n int
    )
    fmt.Scan(&str, &n)
    for i:=range str {
        if i!=n-1{
            fmt.Print(string(str[i]))
        }        
    }
}

//
package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	for _, el := range str {
		if string(el) == "x" || string(el) == "w" {
			fmt.Print(string(el))
			return
		}
	}

	fmt.Print(-1)
}

//
package main

import (
	"fmt"
)

func main() {
	var str, str1 string
	_, _ = fmt.Scan(&str)
	for _, a := range str {
		str1 = string(a) + str1
	}
	if str1 == str {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

//
package main
import "fmt"

func main(){
    var text string
    fmt.Scan(&text)
    
    for i := 0; i < len(text); i++ {
        for j := i+1; j <len(text); j++{
            if text[i] == text[j] {
                fmt.Print(string(text[i]))
                return
            }
        }
    }
}

//
package main
import ("fmt"; "regexp")
func main() {
	var line string
	fmt.Scan(&line)
	re := regexp.MustCompile("^[A-Za-z_]\\w*$")
	if re.MatchString((line)) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

//
package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)   

func main() {
    str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    fmt.Println(strings.Join(strings.Fields(str), " "))
}
//
