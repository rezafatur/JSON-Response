package main

import (
    "JSON_Response/Handler"
    "log"
     "net/http"
)

func main() {
    http.HandleFunc("/nama", Handler.NamaHandler)
    http.HandleFunc("/semuadata", Handler.SemuaDataHandler)
    log.Fatal(http.ListenAndServe(":5050", nil))
}
