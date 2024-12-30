package main
import "fmt"

func main() {
    var n, sum, num int
    _, _ = fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&num)
        if num % 2 == 0 {
            sum++
            fmt.Print(num, " ")
        }
    }
    fmt.Print("\n", sum)
}

//
package main

import (. "fmt")

func main() {
    var size, prev, cur int
    Scan(&size, &prev)
    result := 0
    for i := 1; i < size; i++ {
        Scan(&cur)
        if prev < cur {
            result++
        }
        prev = cur
    }
    Println(result)
}

//
package main

import "fmt"

func main()  {
	var size int
	fmt.Scan(&size)
	array := make([]int, size)
	cnt := 0
	for i:=0; i<size; i++ {
		fmt.Scan(&array[i])
		if array[i] % 3 == 0 && array[i] % 10 == 7 {
			array[i] = -1
			cnt++
		}
	}
	for i:=0; i<size; i++ {
		if array[i] == -1 {
			fmt.Print(cnt, " ")
		} else { fmt.Print(array[i], " ") }
	}
}

//
package main

import "fmt"

func main()  {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	cnt := 1
	fmt.Scan(&numbers[0])
	for i:=1; i<size; i++ {
		fmt.Scan(&numbers[i])
		if numbers[i] > numbers[i - 1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}

//
package main

import (. "fmt")

func main (){
    n , r, tmp := 0, 0, 1
    _,_ = Scan(&n)
    l := make([]int, n)
    for i:=0; i<n; i++{
        _,_ = Scan(&l[i])
    }
    _,_ = Scan(&r)
    
    for inx, val := range l {
        if val>=r{tmp=inx+2}    
    }
    Print(tmp)
    
    
}

//
package main 

import (. "fmt")

func main (){
    n :=0
    _,_ =Scan(&n)
    l :=make([]int, n)
    for i:=0; i<n; i++{
        _,_ = Scan(&l[i])
    
    }
    for inx, val := range l {
        for ind, valm := range l{
            if val+valm==0{
                
                if inx>ind {Println(ind,inx)}
                
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
	var (
		a, c int
		b    [4]int
		d    = [6]int{0, 1, 1, 1, 1, 2}
	)
	_, _ = fmt.Scan(&a)
	for a > 0 {
		a--
		_, _ = fmt.Scan(&c)
		b[c-1]++
	}
	b[0] = max(0, b[0]-b[2])
	fmt.Println(b[3] + b[2] + b[1]/2 + b[0]/4 + d[b[1]%2+b[0]%4])
}

//
package main

import (
	"fmt"
)

func main() {
	var n, h int
	_, _ = fmt.Scan(&n, &h)

	var s = make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&s[i])
	}

	var width = n

	for i := 0; i < n; i++ {
		if s[i] > h {
			width++
		}
	}

	fmt.Println(width)
}

//
package main

import (
    "fmt"
    "strings"
)

func main() {
    var n, m int
    var s string
    res := "#Black&White"
    fmt.Scan(&n, &m)
    for i := n * m; i > 0; i-- {
        fmt.Scan(&s)
        if strings.Contains("CMY", s) {
            res = "#Color"
            break
        }
    }
    fmt.Print(res)
}
