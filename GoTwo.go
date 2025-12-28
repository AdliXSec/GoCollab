package main

import (
	"fmt"
)

func main() {
	var menu int
	var saldo float32
	var log []string

	menu = 0

	for menu != 4 {
		fmt.Println("\nMenu :")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Setor Tunai")
		fmt.Println("3. Tarik Tunai")
		fmt.Println("4. Keluar")

		fmt.Print("Pilih menu : ")
		fmt.Scanln(&menu)

		if menu == 1 {
			fmt.Printf("\nSaldo anda sebesar Rp. %.2f\n", saldo)

			log = append(log, "Cek Saldo [BERHASIL]: Saldo Anda, Rp. "+fmt.Sprintf("%.2f", saldo))
		} else if menu == 2 {
			var setorTunai float32

			fmt.Print("Masukan jumlah setor tunai : ")
			fmt.Scanln(&setorTunai)

			if setorTunai < 0 {
				fmt.Println("\nJumlah setor harus lebih dari 0")

				log = append(log, "Setor Tunai [GAGAL] [Jumlah setor harus lebih dari 0]: Setor Tunai, Rp. "+fmt.Sprintf("%.2f", setorTunai))
			} else {
				saldo += setorTunai
				fmt.Printf("\nSetor tunai sebesar Rp. %.2f berhasil\nSisa saldo kamu sebesar Rp. %.2f\n", setorTunai, saldo)

				log = append(log, "Setor Tunai [BERHASIL]: Setor Tunai, Rp. "+fmt.Sprintf("%.2f", setorTunai)+" Saldo Anda, Rp. "+fmt.Sprintf("%.2f", saldo))
			}
		} else if menu == 3 {
			var tarikTunai float32

			fmt.Print("Masukan jumlah tarik tunai : ")
			fmt.Scanln(&tarikTunai)

			if tarikTunai > saldo {
				fmt.Println("\nSaldo anda tidak mencukupi")

				log = append(log, "Tarik Tunai [GAGAL] [Saldo anda tidak mencukupi]: Tarik Tunai, Rp. "+fmt.Sprintf("%.2f", tarikTunai))
			} else if tarikTunai < 0 {
				fmt.Println("\nJumlah tarik tidak boleh kurang dari 0")

				log = append(log, "Tarik Tunai [GAGAL] [Jumlah tarik tidak boleh kurang dari 0]: Tarik Tunai, Rp. "+fmt.Sprintf("%.2f", tarikTunai))
			} else {
				saldo -= tarikTunai
				fmt.Printf("\nTarik tunai sebesar Rp. %.2f berhasil\nSisa saldo kamu sebesar Rp. %.2f\n", tarikTunai, saldo)

				log = append(log, "Tarik Tunai [BERHASIL]: Tarik Tunai, Rp. "+fmt.Sprintf("%.2f", tarikTunai)+" Saldo Anda, Rp. "+fmt.Sprintf("%.2f", saldo))
			}
		} else if menu == 4 {
			fmt.Println("\nTerimakasih telah menggunakan layanan kami")
			fmt.Println("===== Log Transaksi =====")
			for i, entry := range log {
				fmt.Printf("Log %d: %s\n", i+1, entry)
			}
		}
	}
}
