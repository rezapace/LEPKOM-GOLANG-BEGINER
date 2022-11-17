package main

import "fmt"

const pi = 22 / 7

var jari int

func main(){
    fmt.Print("Masukkan Jari jari lingkaran = ")
    fmt.Scan(&jari)
    luas := pi * (jari * jari)
    fmt.Println("Luas Lingkaran = ", luas)
}