package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	http.HandleFunc("/api/mahasiswa", user)
	fmt.Println("web Berjalan dengan port 5020")
	http.ListenAndServe(":5024", nil)
}

type lepkom struct {
	Nama   string `json:"nama_mahasiswa"`
	Kursus string `json:"kursus_mahasiswa"`
	Foto   string `json:"foto_mahasiswa"`
}

var data_mahasiswa = []lepkom{
	{
		Nama:   "Faldy",
		Kursus: "Golang For Beginer",
		Foto:   "img/gambar1.jpg",
	},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == http.MethodGet {
		result, err := json.Marshal(data_mahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
}
