```
array
var имя_массива [размер_массива]тип
var array [5]int // создаем массив целых чисел
```

```
slice
var slice []int // не задается размер
slice := []int{1,2,3,4,5} // объявили и проинициализировали. slice = [1 2 3 4 5] 
slice := make([]byte, 5) // объявили и проинициализировали. slice = [0 0 0 0 0] 
```

```
for index, value := range collection {
    ...
}

collection - коллекция: массив, слайс, строка
index - индекс текущего элемента коллекции
value - значение текущего элемента коллекции
```

```
var имя [количество строк][количество столбцов]тип
var a [3][4]int // обьявление двумерного массива целых чисел

var number int = a[0][1] // 0 - индекс строки, 1 - индекс столбца
a[0][1] = 8 // элементу 1 строки 2 столбца присвоили значение 8
a := [2][2]string{{"a[0][0]", "a[0][1]"}, {"a[1][0]", "a[1][1]"}}

var s [][]int // объявлен двумерный слайс целых чисел
s := [][]string{{"a[0][0]", "a[0][1]"}, {"a[1][0]", "a[1][1]"}} // объявление с инициализацией двумерного slice-а, размерность строк и столбцов буде выведена автоматически на основании переданных значений
s := make([][]int, 2)
s[0] = make([]int, 5) // 5 столбцов в первой строке
s[1] = make([]int, 10) // 10 столбцов во второй строке
```