package main

import (
	"configs/config"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

func main() {
	cmd := os.Args[1]
	hosts := os.Args[2:]

	config := ssh.ClientConfig{
		User: os.Getenv("LOGNAME"),
		Auth: "password",
	}
	fmt.Println(config)
	fmt.Println("vim-go")
}
