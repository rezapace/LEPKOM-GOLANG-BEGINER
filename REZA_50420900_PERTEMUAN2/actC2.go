package main

import "fmt"

var nilai1, nilai2 float64

func main() {
    defer fmt.Println("---SELESAI---")
    fmt.Print("Masukkan Bilangan 1: ")
    fmt.Scan(&nilai1)
    fmt.Print("Masukkan Bilangan 2: ")
    fmt.Scan(&nilai2)
    hasil := nilai1 / nilai2
    fmt.Printf("Hasil dari nilai1 / nilai2 = %.3f\n", hasil)
}