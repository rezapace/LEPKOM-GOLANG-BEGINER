package main

import "fmt"

func main(){

    var data = map[string]mahasiswa{
        "51420050" : {
            "Rakha Bakhis Attallah",
            "2IA18",
		},
        "10115961" : {
            "Andika Kangen Band",
            "4KA20",
		},
	}	
    var search string
    fmt.Print("Masukkan Npm anda ? ")
    fmt.Scanf("%s", &search)

    var value, ok = data[search]

    if ok {
        fmt.Printf("Nama anda %s \nkelas anda %s", value.Nama, value.Kelas)
    }else{
        fmt.Println("data tidak ditemukan")
    }
}

type mahasiswa struct{
    Nama string
    Kelas string
}