package main

func main(){
    if условие1	{
        блок кода
    } else {
        if условие2	{
            блок кода
        } else {
            if условие3 {
                блок кода
            }
            ...
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
