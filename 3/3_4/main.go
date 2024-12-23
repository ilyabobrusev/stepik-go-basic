package main

import (
	"fmt"
)

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	if a%2 == 0 {
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
	var a int
	fmt.Scan(&a)  
	
	if a%2 != 0 {
		a++
		fmt.Println(a)
	} else {
		a++
		a++
		fmt.Println(a)
	}
}

//
