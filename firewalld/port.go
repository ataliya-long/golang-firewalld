package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	port := flag.String("p", "", "-p 端口号，指定开放某个端口")
	flag.Parse()

	cmd := exec.Command("firewall-cmd", "--zone=public", fmt.Sprintf("--add-port=%s/tcp", *port), "--permanent")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	cmd = exec.Command("firewall-cmd", "--reload")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		return
	}
}

