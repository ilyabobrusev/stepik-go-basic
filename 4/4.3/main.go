package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fact := 1 // переменная, в которой будет хранится итоговое произведение

	for i := 1; i <= n; i++ {
		fact = fact * i // умножаем текущее произведение на число и снова записываем в нее.
	}

	fmt.Println(fact)
}

//

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fact := 1 // переменная, в которой будет хранится итоговое произведение

	for i := a; i <= b; i++ {
		fact = fact * i // умножаем текущее произведение на число и снова записываем в нее.
	}

	fmt.Println(fact)
}
//

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fact := 1

	for i := a; i <= b; i++ {
        if i%10 == 7 || i % 10 == -7 {
            fact = fact * i
        }
		
	}

	fmt.Println(fact)
}

//
package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    fact := 1 
    for i:=2; i<=n; i+=2 {
        fact=fact*i
    }
    fmt.Print(fact)
}

//
