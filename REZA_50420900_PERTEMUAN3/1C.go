package main

import "fmt"

func main(){

    var i int = 1
    for i <= 10 {
        if i%2 == 0 {
            fmt.Println(i, " Adalah Angka Genap")
        } else if i%2 != 0 {
            fmt.Println(i, " Adalah Angka Ganjil")
        } else {
            fmt.Println("Something Wrong")
        }
        i = i + 1
    }
}