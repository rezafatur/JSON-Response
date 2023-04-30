package main

import (
	"fmt"
)

func main() {
    // Deklarasi variabel panjang array
    var length int

    // Input panjang array dari pengguna
    fmt.Print("Masukkan panjang array: ")
    fmt.Scan(&length)

    // Deklarasi variabel array 1 dan array 2
    // Tipe data string dengan panjang sesuai input pengguna
    arr1 := make([]string, length)
	arr2 := make([]string, length)

	fmt.Println("\nInput:")
    // Input array 1 
	fmt.Print("Array 1: ")

	// Looping input array 1
    for i := 0; i < length; i++ {
		fmt.Scan(&arr1[i])
    }

    // Input array 2
	fmt.Print("Array 2: ")

	// Looping input array 2
    for i := 0; i < length; i++ {
		fmt.Scan(&arr2[i])
    }

	fmt.Println("\nOutput:")
    // Looping untuk membandingkan array 1 dan array 2 sesuai indeks
    for i := 0; i < length; i++ {
		// Jika tidak sama, maka keluarkan hasilnya
        if arr1[i] != arr2[i] {
            fmt.Printf("Indeks ke %d berbeda\n", i)
        }
    }
}
