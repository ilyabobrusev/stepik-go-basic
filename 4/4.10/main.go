package main

import (
    "fmt"
)

func main() {
    var n int // количество чисел в последовательности
    _, _ = fmt.Scan(&n)

    var max int    // переменная в которой будет храниться максимальное значение
    var number int // переменная в которую будем считывать значения последовательности

    _, _ = fmt.Scan(&max) // на начальном этапе максимальное значение равно первому значению последовательности

    for i := 1; i < n; i++ { // начинаем с 1 т.к первое значение последовательности считали в max
        _, _ = fmt.Scan(&number) // очередное число последовательности

        if number > max { // сравниваем текущее число с мах
            max = number // если больше, то обновляем мах
        }
    }

    fmt.Println(max) // после просмотра всех чисел последовательности, выводим максимальное значение
}

//
package main

import (
    "fmt"
)

func main() {
    var n int // количество чисел в последовательности
    _, _ = fmt.Scan(&n)

    var max int    // переменная в которой будет храниться максимальное значение
    var number int // переменная в которую будем считывать значения последовательности

    _, _ = fmt.Scan(&min) // на начальном этапе максимальное значение равно первому значению последовательности

    for i := 1; i < n; i++ { // начинаем с 1 т.к первое значение последовательности считали в max
        _, _ = fmt.Scan(&number) // очередное число последовательности

        if number < min { // сравниваем текущее число с мах
            max = number // если больше, то обновляем мах
        }
    }

    fmt.Println(min) // после просмотра всех чисел последовательности, выводим максимальное значение
}

//
package main
import "fmt"

func main(){
    var n, num,min, max int
    fmt.Scan(&n, &num)
    min, max = num, num
    for i := 0; i < n; i++ {
        fmt.Scan(&num)
        if num < min {
            min = num
        }
		if max < num {
            max = num
        }
    }
    fmt.Print(max-min)
}

//
package main
import "fmt"

func main(){
    var n, num, max int
    fmt.Scan(&n, &num)
    max = num
    flag:=" NO"
    for i := 0; i < n-1; i++ {
        fmt.Scan(&num)
        if num > max {
            max = num
        }
        if num <= 30 {
            flag=" YES" 
        }
    }
    fmt.Print(max, flag)
}

//
package main
import "fmt"

func main() {
    var n, s, max int
    _,_ = fmt.Scan(&n)
    min := 30
    condition := false
    for n > 0 {
        _,_ = fmt.Scan(&s)
        if s < min {
            condition = true 
        }
        if s > max {
            max = s
        }
        n--
    }
    if condition == true {
        fmt.Printf("%d %s", max, "YES")    
    } else {
        fmt.Printf("%d %s", max, "NO")
    }
}

//
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	max1 := 0
	max2 := 0
	for n != 0 {
		if n > max1 {
			max2 = max1
			max1 = n
		} else if n > max2 {
			max2 = n
		}
		fmt.Scan(&n)
	}
	fmt.Print(max2)
}
