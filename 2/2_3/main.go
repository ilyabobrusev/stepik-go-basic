package main

import (
	"fmt"
)

func main() {
	var name string
	_, _ = fmt.Scan(&name)
	fmt.Println("Привет,", name)
}

#

package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    _ = scanner.Scan()
    name := scanner.Text()
    fmt.Println(name, "- лучшая книга!")
}

#

package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
	// step input
    _ = scanner.Scan()
    var1 := scanner.Text()
    _ = scanner.Scan()
    var2 := scanner.Text()
    _ = scanner.Scan()
    var3 := scanner.Text()
    // step print
    fmt.Println(var1)
    fmt.Println(var2)
    fmt.Println(var3)
}

#

package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    // step input
    _ = scanner.Scan()
    var1 := scanner.Text()
    _ = scanner.Scan()
    var2 := scanner.Text()
    _ = scanner.Scan()
    var3 := scanner.Text()
    // step print
    fmt.Println(var3)
    fmt.Println(var2)
    fmt.Println(var1)
}

#

package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    // step input
    _ = scanner.Scan()
    spl := scanner.Text()
    _ = scanner.Scan()
    var1 := scanner.Text()
    _ = scanner.Scan()
    var2 := scanner.Text()
    _ = scanner.Scan()
    var3 := scanner.Text()
    // step print
    fmt.Print(var1,spl,var2,spl,var3)
}
