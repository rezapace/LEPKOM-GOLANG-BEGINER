package main

import "fmt"

func main(){
    var uts, uas, total_nilai float64
    fmt.Print("Masukkan Nilai UTS : ")
    fmt.Scan(&uts)
    fmt.Print("Masukkan Nilai UAS : ")
    fmt.Scan(&uas)


    fmt.Print("\n Nilai UTS = ", uts)
    fmt.Print("\n Nilai UAS = ", uas)

    total_nilai = (uts * 0.3) + (uas * 0.7)
    fmt.Print("\n Total Nilai = ", total_nilai)

    if total_nilai <= 20 {
        fmt.Print("\n Grade E")
    } else if total_nilai <= 40 {
        fmt.Print("\n Grade D")
    } else if total_nilai <= 60 {
        fmt.Print("\n Grade C") 
    } else if total_nilai <= 80 {
        fmt.Print("\n Grade B")
    } else {
        fmt.Print("\n Grade A")
        }
    }