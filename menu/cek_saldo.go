package menu

import (
	"fmt"

	"github.com/Agilistikmal/atm/database"
	"github.com/Agilistikmal/atm/utils"
)

func CekSaldo() {
	fmt.Println("| --------------- CEK SALDO ---------------")
	fmt.Printf("| Saldo anda Rp%d", database.Saldo)
	fmt.Println("")
	utils.Pause()
}

