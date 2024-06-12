package tools

import (
	"bytes"
	"log"
	"os/exec"
)

func HostNameChanger(newName string, password string) string {

	cmd := exec.Command("sudo", "-S", "hostnamectl", "set-hostname", newName)

	var stdin bytes.Buffer
	stdin.Write([]byte(password + "\n"))
	cmd.Stdin = &stdin

	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	cmd2 := exec.Command("hostname")

	output, err := cmd2.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}
