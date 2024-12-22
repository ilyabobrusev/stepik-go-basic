package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 {
		fmt.Print("Положительное")
	}
	if n < 0 {
		fmt.Print("Отрицательное")
	}
	if n == 0 {
		fmt.Print("Ноль")
	}
}

//

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 0 {
		fmt.Print("Положительное")
	} else {
		if n < 0 {
			fmt.Print("Отрицательное")
		} else {
			fmt.Print("Ноль")
		}
	}
}

//
package main

import "fmt"

func main() {
    var grade int
    fmt.Scan(&grade)
    if grade >= 90 {
        fmt.Print(5)
    } else {
        if grade >= 80 {
            fmt.Print(4)
        } else {
            if grade >= 70 {
                fmt.Print(3)
            } else {
                if grade >= 60 {
                    fmt.Print(2)
                } else {
                    if grade >= 0 {
                        fmt.Print(1)
                    }
                }
            }
        }
    }
}

//
package main

import "fmt"

func main() {
	var n int
	_,_ =fmt.Scan(&n)

	if n >= -3 && n <= 1 {
		fmt.Print("YES")
	} else {
		if n >= 5 && n <= 9 {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	}
}

//
