package main

import (
    "fmt"
)

func main() {
    for i := 1; i <= 6; i++ {
        fmt.Print("5 ")
    }
}

//
package main

import (
    "fmt"
)

func main() {
    for k := 1; k <= 3; k++ { // выполняется 3 раза
        for i := 1; i <= 6; i++ {
            fmt.Print("5 ") //вывод одной строки
        }
        fmt.Println() // переводим курсор на следующую строку. После того как вывели всю строку!
    }
}

//
package main

import (
	"fmt"
)

func main() {
	a, b, k := 0, 0, 0
	_, _ = fmt.Scan(&a, &b, &k)

	for a <= b {
		cnt := 0
		for i := 1; i <= a; i++ {
			if a%i == 0 {
				cnt++
			}
		}
		if cnt >= k {
			fmt.Print(a, " ")
		}
		a++
	}
}

//
package main

import "fmt"

func main() {
    var n, counter int
    
    fmt.Scanln(&n)
    
    for i := 7; i <= n; i++ {
        for j := i; j > 0; {
            if j % 10 == 7 {
                counter ++
            }
            j /= 10
        }
    }
    
    fmt.Println(counter)
}

//
package main

import (
    "fmt"
)

func main() {
    var a  int
    fmt.Scan(&a)
    sum:= 0
    maxsum:= 0
    maxn:=0
    for j := 1; j <= a; j++ { 
        n:= j
        for n > 0 {
            dig:= n%10
            sum=sum+dig
            n = n/10  
        }    
        if sum > maxsum {
            maxsum = sum
            maxn=j
        }
        sum = 0
    }
    fmt.Println(maxn, maxsum)
}

//
package main
import "fmt"

func main(){
    var num int
    fmt.Scan(&num)
    
    for i := 2; i <= num; i++ {
        for num % i == 0 {
            fmt.Print(i, " ")
            num = num/i
        }
    }
}

//
