package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Agilistikmal/atm/database"
	"github.com/Agilistikmal/atm/utils"
)

func Transfer() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("| --------------- TRANSFER ---------------")
	fmt.Println("| Saldo anda => Rp", database.Saldo)
	fmt.Println("| *Ketik 0 untuk keluar.")
	fmt.Print("| Masukkan nomor rekening penerima => ")
	scanner.Scan()
	rekening_tujuan := scanner.Text()

	if rekening_tujuan == "0"{
		Menu()
		return
	}

	fmt.Print("| Masukkan jumlah yang ingin di transfer => ")
	scanner.Scan()
	jumlah, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("! Error: Jumlah yang dimasukkan bukan angka")
	}

	if jumlah > database.Saldo {
		fmt.Println("! Error: Saldo anda tidak mencukupi ( Rp", database.Saldo, ")")
	} else {
		if jumlah < 10_000 {
			fmt.Println("! Error: Minimal transfer adalah Rp10000")
		} else {
			fmt.Println("# Berhasil melakukan transfer Rp", jumlah, "ke rekening", rekening_tujuan)
		}
	}

	utils.Pause()
	utils.Clear()
	Transfer()
}