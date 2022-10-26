package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Agilistikmal/atm/menu"
)

func main() {
	run := true
	for run {
		menu.Menu()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("# Pilih menu => ")
		scanner.Scan()
		choose := scanner.Text()

		switch choose {
		case "1":
			menu.CekSaldo()
			break
		case "2":
			menu.Transfer()
			break
		case "3":
			menu.TarikTunai()
			break
		case "0":
			run = false
		}

		// run = false
	}
}