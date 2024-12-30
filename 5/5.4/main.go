package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := n - 1 - i; j > n-2-i; j-- {
			nums[i][j] = 1
			for d := j + 1; d < n; d++ {
				nums[i][d] = 2
			}
		}
	}
	for i := range nums {
		for j := range nums[i] {
			fmt.Print(nums[i][j], " ")
		}
		fmt.Println()
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
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != matrix[j][i] {
				fmt.Print("NO")
				return
			}
		}
	}

	fmt.Print("YES")
}

//
package main
import "fmt"
func main(){
    var n, m int
    _,_=fmt.Scan(&n,&m)
    
    for i:=1;i<=n;i++{
        for j:=1;j<=m;j++{
            fmt.Print(i*j," ")
        }
        fmt.Println()
    }

}

//
package main
import ("fmt")
func main() {
	var n int
	fmt.Scan(&n)
    var m int
    fmt.Scan(&m)    
    matrix := make([][] int, n)   
    for i:= 0; i <n; i++{
        matrix[i] = make([]int, m)
        for j := 0; j < m; j++ {    
            if j == 0 || i == 0 {
                matrix [i][j] = 1
            }else {
            matrix[i][j] = matrix[i-1][j]+matrix[i][j-1]
            }
        fmt.Print(matrix[i][j], " ")
    }
        fmt.Println(" ")
}
}

//
package main

import "fmt"

func main() {
	var strlen, columnlen int
	_, _ = fmt.Scan(&strlen, &columnlen)

	nums := make([][]int, strlen)
	sum, lastSum, index := 0, 0, 0

	for i := 0; i < strlen; i++ {
		nums[i] = make([]int, columnlen)

		for j := 0; j < columnlen; j++ {
			_, _ = fmt.Scan(&nums[i][j])
			lastSum += nums[i][j]
		}
		if lastSum > sum {
			sum = lastSum
			index = i
		}
		lastSum = 0
	}

	fmt.Println(sum, "\n", index)
}

//
