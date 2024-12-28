package main

import (
    "fmt"
)

func main() {
    for n := 1; n <= 10; n++ {
        if n%2 == 0 {
            continue // переход к следующей итерации цикла
        }
        if n == 7 {
            break // прерывание цикла
        }

        fmt.Println(n)
    }
}

//
package main

import "fmt"

func main()  {
	var n  int
	fmt.Scan(&n)
	for d := n/2; d < n; d-- {
		if n % d == 0 {
			fmt.Println(d)
			break
		}
	}
}

//
package main

import "fmt"

func main()  {
	var n  int
	for fmt.Scan(&n); n <= 100; fmt.Scan(&n) {
		if n < 10 {continue}
		fmt.Println(n)
	}
}

//
package main

import "fmt"

func main() {
    var n,c,d int
    fmt.Scan(&n,&c,&d)
    
    for i:=1 ; i<=n ; i++{
        if i%c==0 && i%d!=0{
            fmt.Println(i)
            break
        }
    }
}

//
package main

import "fmt"

func main()  {
	var s1, sn string
	fmt.Scan(&s1)
	cnt := 1
	for fmt.Scan(&sn); s1 != sn; fmt.Scan(&sn) {
		cnt ++
	}
	fmt.Println(cnt)
}
