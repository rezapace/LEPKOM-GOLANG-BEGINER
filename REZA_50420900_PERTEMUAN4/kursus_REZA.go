package main

import "fmt"

func main()  {
    var kursus = []string{"dbms", "server os", "networking", "web", "desktop", "erp"}
    fmt.Print("Isi dari kursus awal adalah : ", kursus)
    fmt.Print("\nPanjang kursus ", len(kursus), " kapasitas ", cap(kursus), "\n")

    var ambil = kursus[0:3]
    fmt.Print("\nIsi dari kursus 2 adalah : ", ambil)
    fmt.Print("\nPanjang kursus ", len(ambil), " kapasitas ", cap(ambil), "\n")

    var tambah = append(ambil, "Manajemen Informatika")
    fmt.Print("\nIsi dari kursus 3 adalah : ", tambah)
    fmt.Print("\nPanjang kursus ", len(tambah), " kapasitas ", cap(tambah), "\n")
}