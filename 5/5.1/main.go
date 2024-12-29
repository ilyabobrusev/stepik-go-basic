package main

import "fmt"

func main() {
    var numbers [5]int
    numbers[0] = 5 // записать в первый элемент массива 5
    numbers[1] = 2
    numbers[4] = 3
//    numbers[5] = 2          // ошибка, индекс вне границ массива
    fmt.Println(numbers[0]) // вывести первый элемент массива
}

//
package main

import "fmt"

func main() {
    var array [5]int
    fmt.Println(len(array))
}

//
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice) // вывод: [1 2 3 4 5]
    slice = append(slice, 6, 7, 8)
    fmt.Println(slice) // вывод: [1 2 3 4 5 6 7 8]
}

//
package main

import "fmt"

func main() {
    slice := make([]int, 5) // будет создан слайс из 5 элементов со значениями по умолчанию 0 для int-а
    fmt.Println(slice) // вывод: [0 0 0 0 0]
}

//

package main

import "fmt"

func main() {
    var size int
    fmt.Scan(&size)                     // считываем сколько чисел будем вводить с клавиатуры
    numbers := make([]int, size)        // заводим новый slice целого типа размерностью кол-ва чисел, которые будем вводить с клавиатуры
    for i := 0; i < len(numbers); i++ { // получаем длину slice-а и проходим по нему циклом
        fmt.Scan(&numbers[i]) // считываем и сразу записываем значение в i-ый индекс slice-а
    }
    fmt.Println(numbers)
}

//
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(numbers); i++ {
        fmt.Print(numbers[i], " ") // или fmt.Println(numbers[i])
    }
}

//
package main

import "fmt"

func main() {
    numbers := make([]int, 5)
    numbers[0] = 1
    numbers[1] = 2
    numbers[2] = 3
    numbers[3] = 4
    numbers[4] = 5
    fmt.Println(numbers)
}
//

package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scan(&size)
	slice := make([]int, size)
	for i := 0; i < len(slice); i++  {
		fmt.Scan(&slice[i])
		if i%3 == 0 {
			fmt.Print(slice[i], " ")
		}
	}
}

//
package main
import "fmt"
func main () {
    var n int
    _,_=fmt.Scan(&n)
    arr:=make([]int,n)
    for i:=0 ; i<len(arr) ; i++ {
        _,_=fmt.Scan(&arr[i])
        if arr[i]%3==0 {
            fmt.Print (arr[i]," ")
        }
    }
}

//
package main
import "fmt"

func main(){
    var size,count int
    fmt.Scan(&size)
    array := make([]int, size)
    fmt.Scan(&array[0])
    for i := 1; i < len(array); i++ {
        fmt.Scan(&array[i])
        if array[i-1] < array[i] {
            count ++
        }
    }
    fmt.Print(count)
}

//
package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)

	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}

	for i := 0; i < len(numbers)-1; i += 2 {
		numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
}

//
package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	count := 0
	var array = make([]int, size)
	fmt.Scan(&array[0])
	for i := 1; i < len(array); i++ {
		fmt.Scan(&array[i])
		if array[i]*array[i-1] >= 0 {
			count++
		}
	}
	if count > 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

//
package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	for i:=1; i<size; i++ {fmt.Scan(&numbers[i])}
	fmt.Scan(&numbers[0])
	for i:=0;i<size;i++ {fmt.Print(numbers[i], " ")}
}