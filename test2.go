package main

import "fmt"

func main() {
	var input int
	var saldo float32
	var log []string

	for input != 4 {
		fmt.Printf("=== Bank TelU ===\n")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Setor Uang")
		fmt.Println("3. Tarik Uang")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&input)

		switch {
		case input == 1:
			fmt.Printf("Saldo Anda: Rp %.2f\n\n\n", saldo)

			// LOG
			out := fmt.Sprintf("Cek Saldo: [BERHASIL] Saldo Anda, Rp %.0f", saldo)
			log = append(log, out)
		case input == 2:
			var setor float32
			fmt.Print("Masukkan jumlah uang yang ingin disetor: ")
			fmt.Scan(&setor)

			if setor < 0 {
				fmt.Printf("Jumlah setoran harus lebih dari 0!\n\n\n")

				// LOG
				out := fmt.Sprintf("Setor Uang: [GAGAL] Rp %.0f (Jumlah tidak valid) ", setor)
				log = append(log, out)
			} else {
				saldo += setor
				fmt.Printf("Saldo Anda sekarang: Rp %.2f\n\n\n", saldo)

				// LOG
				out := fmt.Sprintf("Setor Uang: [BERHASIL] Rp %.0f Saldo Anda, Rp %.0f", setor, saldo)
				log = append(log, out)
			}
		case input == 3:
			var tarik float32
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: ")
			fmt.Scan(&tarik)

			if tarik < 0 {
				fmt.Printf("Jumlah penarikan harus lebih dari 0!\n\n\n")

				// LOG
				out := fmt.Sprintf("Tarik Uang: [GAGAL] Rp %.0f (Jumlah tidak valid)", tarik)
				log = append(log, out)
			} else if tarik > saldo {
				fmt.Printf("Saldo Anda tidak mencukupi!\n\n\n")

				// LOG
				out := fmt.Sprintf("Tarik Uang: [GAGAL] Rp %.0f (Saldo tidak mencukupi)", tarik)
				log = append(log, out)
			} else {
				saldo -= tarik
				fmt.Printf("Saldo Anda sekarang: Rp %.2f\n\n\n", saldo)

				// LOG
				out := fmt.Sprintf("Tarik Uang: [BERHASIL] Rp %.0f Saldo Anda, Rp %.0f", tarik, saldo)
				log = append(log, out)
			}
		case input == 4:
			fmt.Printf("Terima kasih\n\n")

			// LOG
			fmt.Print("=== Log Transaksi ===\n")
			for i := 0; i < len(log); i++ {
				fmt.Println(log[i])
			}
		default:
			fmt.Printf("Pilihan tidak valid\n\n\n")
		}
	}
}
