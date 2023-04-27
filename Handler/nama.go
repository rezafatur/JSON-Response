package Handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Nama string `json:"Nama"`
    NIM string `json:"NIM"`
    Alamat string `json:"Alamat"`
}

var data []Data

func NamaHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var d Data
        d.Nama = r.FormValue("Nama")
        d.NIM = r.FormValue("NIM")
        d.Alamat = r.FormValue("Alamat")

		// Data kosong tidak bisa di-POST
        if d.Nama == "" || d.NIM == "" || d.Alamat == "" {
            http.Error(w, "Data yang diterima tidak lengkap", http.StatusBadRequest)
            return
        }

		// NIM sama tidak bisa di-POST
        for _, existingData := range data {
            if existingData.NIM == d.NIM {
                http.Error(w, "NIM sudah ada dalam data", http.StatusBadRequest)
                return
            }
        }

		// Add data
        data = append(data, d)
        jsonData, err := json.Marshal(d)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, string(jsonData))
    } else {
        http.Error(w, "Metode yang diperbolehkan hanya POST", http.StatusMethodNotAllowed)
    }
}
