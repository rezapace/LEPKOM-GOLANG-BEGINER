package main

import "fmt"

var nilai1, nilai2, nilai3, nilai4, nilai5 float64

func main() {
   fmt.Print("masukkan Nilai 1 = ")
   fmt.Scan(&nilai1)
   fmt.Print("masukkan Nilai 2 = ")
   fmt.Scan(&nilai2)
   fmt.Print("masukkan Nilai 3 = ")
   fmt.Scan(&nilai3)
   fmt.Print("masukkan Nilai 4 = ")
   fmt.Scan(&nilai4)
   fmt.Print("masukkan Nilai 5 = ")
   fmt.Scan(&nilai5)
   hasil := (nilai1 * nilai2) + nilai3 - (nilai4 / nilai5)
   fmt.Printf("Hasil = (%.3f * %.3f)+%.3f-(%.3f / %.3f) = %.3f \n", nilai1, nilai2, nilai3, nilai4, nilai5, hasil)
}