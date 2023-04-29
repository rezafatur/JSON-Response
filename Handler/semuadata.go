package Handler

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func SemuaDataHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        jsonData, err := json.Marshal(data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, string(jsonData))
    } else {
        http.Error(w, "Metode yang diperbolehkan hanya GET", http.StatusMethodNotAllowed)
    }
}
