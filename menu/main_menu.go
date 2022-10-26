package menu

import (
	"fmt"

	"github.com/Agilistikmal/atm/utils"
)

func Menu() {
	utils.Clear()
	fmt.Println("| --------------- ATM ---------------")
	fmt.Println("| Selamat datang di mesin ATM.")
	fmt.Println("| Silahkan pilih menu dibawah ini.")
	fmt.Println("| 1) Cek Saldo")
	fmt.Println("| 2) Transfer")
	fmt.Println("| 3) Tarik Tunai")
	fmt.Println("| ")
	fmt.Println("| 0) Keluar")
	fmt.Println("| --------------- ATM ---------------")
}