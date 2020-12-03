package main

import (
	"log"
	"os/exec"
	_ "strings"
)

func main() {
	proc := exec.Command("go", "list", "...")
	raw, _ := proc.Output()

	s := string(raw)

	var i int = 0
	var j int = 0
	for i = 0; i < len(s); i++ {
		//fare che conta ogni pacchetto
		if raw[i] == 10 {
			j++
		}
	}

	log.Println(j)
}
