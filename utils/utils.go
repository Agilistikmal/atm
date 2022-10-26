package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Pause() {
	fmt.Print("\n? Klik tombol apa saja untuk melanjutkan...")
  bufio.NewReader(os.Stdin).ReadBytes('\n') 
}