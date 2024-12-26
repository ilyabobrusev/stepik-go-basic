package main

import "fmt"

func main() {
	sum := 0
	for i := 15; i > 1; i-- {
		sum = sum + i
	}
	fmt.Println(sum)
}

//
package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    count := 0
    for i := 1; i <= n; i++ {
        if i%3 == 0 && i%7 != 0 {
            count++
        }
    }

    fmt.Println(count)
}

//
