package main

import (
	"fmt"
	"math/rand"
)

func main() {
	angkabenar := rand.Intn(100)
	jumlah := 0

	for {
		var angka int
		fmt.Print("Masukan angka yang di tebak : ")
		fmt.Scanln(&angka)
		if angka == angkabenar {
			fmt.Println("\nSelamat anda benar")
			jumlah++
			fmt.Printf("Anda memerlukan %d kali kesempatan", jumlah)
			break
		} else {
			fmt.Println("\nMaaf anda salah")
			if angka > angkabenar {
				fmt.Println("Tebakan anda terlalu besar\n")
			} else if angka < angkabenar {
				fmt.Println("Tebakan anda terlalu kecil\n")
			}
			jumlah++
		}
	}
}
