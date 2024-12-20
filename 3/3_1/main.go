package main

import "fmt"

func main() {
	fmt.Println("Выйти из дома")
	fmt.Println("Дойти до магазина")
	fmt.Println("Хлеб есть? Введите Есть или Нет")
	
	var answer string
	fmt.Scan(&answer)
	
	if answer == "Есть" {
		fmt.Println("Купить хлеб")
	}

	fmt.Println("Дойти до дома")
	fmt.Println("Зайти в дом")
}

//
package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x > 0 {
		fmt.Println("Положительное")
	} else {
		fmt.Println("Отрицательное")
	}
}

//
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

//
package main

import "fmt"
 
func main() {
	a := 2
	b := 3
	c := 2
	fmt.Println(a > b) // false
	fmt.Println(a < b) // true
	fmt.Println(a >= b) // false
	fmt.Println(a >= c) // true
	fmt.Println(a == b) // false
	fmt.Println(a == c) // true
	fmt.Println(a != b) // true
}

//
package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Print(1)

	if x != 5 {
		fmt.Print(2)
	}

	fmt.Print(3)
}

//
package main

import "fmt"

func main(){
    var x int
    fmt.Scan(&x)

    fmt.Print(1)

    if x % 10 <= 7{
        fmt.Print(2)
    } else {
        fmt.Print(3)
    }
    fmt.Print(4)
}

//
package main

import "fmt"

func main(){
    var a, b int
    fmt.Scan(&a, &b)

    if a > b{
        fmt.Print(a)
    } else if b > a {
        fmt.Print(b)   
    } else {
        fmt.Print(a)
    }
}

//
package main

import "fmt"

func main(){
    var a, b string
    fmt.Scan(&a, &b)

    if a == b{
        fmt.Println("Пароль принят")  
    } else {
        fmt.Println("Пароль не принят")
    }
}

//
package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)

    if a >= 12{
        fmt.Println("Доступ разрешен")  
    } else {
        fmt.Println("Доступ запрещен")
    }
}

//
package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)

    if a % 2 == 0 {
        fmt.Println("YES")  
    } else {
        fmt.Println("NO")
    }
}

//
package main

import "fmt"

func main() {
    var a,b int
    fmt.Scan(&a,&b)
    
    if a%b==0 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

//
package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)

    if a < 0{
        fmt.Print(-1)
    } else if a == 0 {
        fmt.Print(0)   
    } else {
        fmt.Print(1)
    }
}
