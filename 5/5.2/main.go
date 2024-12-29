package main

import "fmt"

func main() {
    var count int // сколько чисел будут вводиться на следующей строке
    fmt.Scan(&count)
    numbers := make([]int, count)
    for i := 0; i < count; i++ {
        fmt.Scan(&numbers[i]) // инициализируем коллекцию значениями
    }

    // изначально индекс максимального и максимальное значение равны первому элементу
    indexMax := 0                // здесь храним индекс максимального, равен первому элементу коллекции
    max := numbers[indexMax]     // здесь храним максимальное, равен первому элементу коллекции
    for i := 1; i < count; i++ { // начинаем перебирать коллекцию, начинаем с 1 потому как indexMax и max проинициализировали первым элементом коллекции
        if numbers[i] > max { // если текущий элемент коллекции больше максимального, то обновляем максимальное и индекс максимального
            max = numbers[i]
            indexMax = i
        }
    }
    fmt.Print(max, "\n", indexMax, "\n")
}

//
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // для корректного ввода строки с пробелами воспользуемся пакетом bufio
    myscanner := bufio.NewScanner(os.Stdin)
    myscanner.Scan()
    input := myscanner.Text()

    stringNumbers := strings.Split(input, " ") // функция Split из пакета strings, разбивает строку из переменной str на коллекцию строк. Разделителем будет пробел
    var numbers []int
    for _, stringNumber := range stringNumbers {
        number, _ := strconv.Atoi(stringNumber) // преобразуем строку в число
        numbers = append(numbers, number)       // добавляем значения в коллекцию
    }
    length := len(numbers)
    for i := 0; i < length/2; i++ {
        tmp := numbers[i]
        numbers[i] = numbers[length-1-i]
        numbers[length-1-i] = tmp
    }
    for _, number := range numbers {
        fmt.Print(number, " ")
    }
}

//
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	minInd := 0

	for index := range numbers {
		fmt.Scan(&numbers[index])
		if numbers[index] < numbers[minInd] {
			minInd = index
		}
	}
	fmt.Println(minInd)
}

//
package main

import (
	"fmt"
)

func main() {
	n, minValue := 0, 10001
	_, _ = fmt.Scan(&n)
	list := make([]int, n)
	for index := range list {
		_, _ = fmt.Scan(&list[index])
		if list[index] < minValue {
			minValue = list[index]
		}
	}
	for index := range list {
		fmt.Print(list[index]-minValue, " ")
	}
}

//
package main
import "fmt"

func main() {
    var a, min, max int
    fmt.Scan(&a)
    b := make([]int, a)
    for i := 0; i < a; i++{
        fmt.Scan(&b[i])
        if b[i] < b[min]{ min = i
        }else if b[i] > b[max]{ max = i}
    } 
    fmt.Print(max - min)
}

//
package main
import "fmt"

func main() {
a := []int{7,9,8,1,2,3,3,10,8,6}
s := 0
for i := 1; i < 10; i++{
    if a[i - 1] < a[i]{ 
        a[i] = a[i - 1] + 1
        s = s + 1
    }  
}
fmt.Print(s)
}

//
package main

import ("fmt")

func main() {
    var n int
    minValue, minIndex, maxValue, maxIndex := 10001, -1, -10001, -1

    fmt.Scan(&n)
    array := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&array[i])
        if array[i] < minValue {
            minValue, minIndex = array[i], i
        }
        if array[i] >= maxValue {
            maxValue, maxIndex = array[i], i
        }
    }
    array[minIndex], array[maxIndex] = maxValue, minValue
    for i := 0; i < n; i++ {
        fmt.Print(array[i], " ")
    }
}

//
package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	result := "YES"
	for i := 0; i < size; i++ { fmt.Scan(&numbers[i]) }
	for i := 0; i < size / 2; i++ {
		if numbers[i] != numbers[size - 1 - i] {
			result = "NO"
			break
		}
	}
	fmt.Println(result)
}
