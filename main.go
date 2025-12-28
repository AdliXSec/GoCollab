package main

import (
	"bufio"
	"fmt"
	"os"

	"judol-detector/detector"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan URL: ")
	scanner.Scan()

	url := scanner.Text()

	if detector.IsJudol(url) {
		fmt.Println("🚨 TERDETEKSI LINK JUDOL")
	} else {
		fmt.Println("✅ URL AMAN")
	}
}
