package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Println(num * num)
}

//

package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Println(num / 1000)
}

//

package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	num = num * num * num
	fmt.Println(num * num)

}

//

package main

import "fmt"

func main() {
	var num1, num2, num3 int
	_, _ = fmt.Scan(&num1, &num2, &num3)
	fmt.Println(num1 * num2 * num3)

}

//

package main

import "fmt"

func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)
	fmt.Println(k / n)

}

//

package main

import "fmt"

func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)
	fmt.Println(k % n)

}

//

package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Println(num)
	num++
	fmt.Println(num)
	num++
	fmt.Println(num)
}

//

package main

import "fmt"

func main() {
	var t, c, p, h int
	_, _ = fmt.Scan(&t, &c, &p, &h)
	fmt.Println(3 * (t + c + p + h))
}

//

package main

import "fmt"

func main() {
	var r, k, n int
	fmt.Scan(&r, &k, &n)
	fmt.Println(r*n+((k*n)/100), (k*n)%100)
}

//

package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)
	fmt.Println(m, "мин - это", m/60, "час", m%60, "минут.")
}

//

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println("Следующее за числом", num, "число:", num+1)
	fmt.Println("Для числа", num, "предыдущее число:", num-1)
}
