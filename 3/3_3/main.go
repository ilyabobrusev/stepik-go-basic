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

package main

import "fmt"

func main() {
	var n int
	_,_ =fmt.Scan(&n)

	if n < 60 {
		fmt.Print("Легкий вес")
	} else {
		if n < 64 {
			fmt.Print("Первый полусредний вес")
		} else {
			fmt.Print("Полусредний вес")
		}
	}
}

//
package main

import "fmt"

func main() {
    var n1, n2, n3 int
    _,_ = fmt.Scan(&n1, &n2, &n3)
    if n1 == n2  && n1 == n3 {
        fmt.Print(3)
    } else {
        if n1 == n2 || n1 == n3 || n2 == n3 {
            fmt.Print(2)
        } else {
            fmt.Print(0)        
        }
    }
}

//
package main

import ("fmt")

func main(){
    var a, b float64
    
    fmt.Scan(&a, &b)

    if a > 0 {
        if b > 0 {
            fmt.Println(1)
        } else {
            fmt.Println(4)
        }
    } else {
        if b > 0 {
            fmt.Println(2)
        } else {
            fmt.Println(3)
        }
    }
}

//

package main

import "fmt"

func main() {
	var x, y int
	var s string
	fmt.Scan(&x, &y)
	fmt.Scan(&s)

	if y == 0 && s == "/" {
		fmt.Println("На ноль делить нельзя!")
	} else {
		if s == "/" {
			fmt.Println(float64(x) / float64(y))
		} else {
			if s == "+" {
				fmt.Println(x + y)
			} else {
				if s == "-" {
					fmt.Println(x - y)
				} else {
					if s == "*" {
						fmt.Println(x * y)
					} else {
						fmt.Println("Неверная операция")
					}
				}
			}
		}
	} 
}

//
package main

import (
	"fmt"
	"math"	
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
    
	D := math.Pow(b, 2) - 4*a*c

	if D >= 0 {
		if D == 0 {
			fmt.Println(-b/(2*a))
		} else if D > 0 {
			x1 := (-b + math.Sqrt(D))/(2*a)
			x2 := (-b - math.Sqrt(D))/(2*a)
			if x1 > x2 {
				fmt.Println(x2)
				fmt.Println(x1)
			} else {
				fmt.Println(x1)
				fmt.Println(x2)
			}
		}
	}
}

//

package main

import "fmt"

func main() {
    var action string
    fmt.Scan(&action)
    switch action {
    case "Прыгать":
        fmt.Println("Мне нравится прыгать.")
    case "Плавать":
        fmt.Println("Я люблю плавать.")
    case "Летать":
        fmt.Println("Хотел бы я научиться летать.")
    default:
        fmt.Println("Могу только ходить.")
    }
}

//
package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

    switch a {
    case "1":
        fmt.Println("31")
    case "2":
        fmt.Println("29")
    case "3":
        fmt.Println("31")
	case "4":
        fmt.Println("30")
	case "5":
        fmt.Println("31")
	case "6":
        fmt.Println("30")
	case "7":
        fmt.Println("31")
	case "8":
        fmt.Println("31")
	case "9":
        fmt.Println("30")
	case "10":
        fmt.Println("31")
	case "11":
        fmt.Println("30")
	case "12":
        fmt.Println("31")
    }
}

//
package main 
 
import ( 
    "fmt"
) 
 
func main() { 
 
    var month string
    fmt.Scanln(&month) 

    switch month { 
        case "12", "1", "2": 
            fmt.Println("Зима") 
        case "3", "4", "5": 
            fmt.Println("Весна") 
        case "6", "7", "8": 
            fmt.Println("Лето") 
        case "9", "10", "11": 
            fmt.Println("Осень") 
    } 
}

//
package main 
 
import ( 
    "fmt"
) 
 
func main() { 

	a := 56
	b := 7 
	k := 6
	a = a / 7 - b
	var c int 
	if a > b {   
		c = a - k * b
	} else {   
		c = a + k * b
	}
	fmt.Print(c)
}