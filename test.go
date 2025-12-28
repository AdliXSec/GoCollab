package main

import (
	"fmt"
)

func main() {
	var log []string // deklarasi slice log

	// Menambahkan elemen ke dalam slice log
	log = append(log, "Started application")
	log = append(log, "Connected to database")
	log = append(log, "Application stopped")

	// Menampilkan semua log
	// for i, entry := range log {
	// 	fmt.Printf("Log %d: %s\n", i+1, entry)
	// }
	fmt.Println(log)
}
