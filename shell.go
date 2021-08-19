package main

import (
	"os/exec"
	"log"
)

func main() {
	cmd := exec.Command("./start.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
