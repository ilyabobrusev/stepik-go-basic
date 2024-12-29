package main

import "fmt"

func main() {
    var count int
    fmt.Scan(&count)
    numbers := make([]int, count)
    for i := 0; i < count; i++ {
        fmt.Scan(&numbers[i])
    }

    for i := 0; i < 10; i++ { // внешний цикл, в котором перебираем числа, кол-во которых нам необходоимо подсчитать
        cnt := 0                         // переменная, которую будем использовать для подсчета кол-ва элементов равных i
        for _, number := range numbers { // внутренний цикл, в котором перебираем элементы, которые ввели с клавиатуры
            if number == i { // если i (значение из внешнего цикла) равен элементу коллекции введенной с клавиатуры, то увеличиваем счетчик cnt
                cnt++
            }
        }
        // в этом моменте мы уже сравнили все элементы массива с i
        fmt.Println("Количество чисел", i, "равно", cnt)
    }
}

//
package main

import "fmt"

func main() {
    var count int
    fmt.Scan(&count)
    numbers := make([]int, count)
    
    for i := 0; i < count; i++ {
        fmt.Scan(&numbers[i])
    }

	
	for i:=0;i<count;i++{
   		 for index,value := range numbers{
			if value == numbers[i] && index !=i{
				fmt.Print("YES")
				return
			}
	
		}
	}
	fmt.Print("NO")
}

//
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	slice := make([]int, n)
	
	for i := 0; i < len(slice); i++ {
		fmt.Scan(&slice[i])
	}
	for i := 0; i < len(slice); i++ {
		kol := 0
		for _, number := range slice {
			if number == slice[i] {
				kol++
			}
		}
		if kol == 1 {
			fmt.Print(slice[i], " ")
		} 
	}
}

//
package main

import "fmt"

func main() {
    var count int
    fmt.Scan(&count)
    numbers := make([]int, count)
    for i := 0; i < count; i++ {
        fmt.Scan(&numbers[i])
    }

    cnt := 0
    numbersSize := len(numbers)
    for i, number := range numbers {
        for j := i + 1; j < numbersSize; j++ {
            if number == numbers[j] {
                cnt++
            }
        }
    }
    fmt.Println(cnt)
}

