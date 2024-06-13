package tools

import (
	"fmt"
	"github.com/GrosbergKirr/Server_hostname/internal/logger"
	"os/exec"
)

func HostNameChanger(newName string, password string, ok chan string) error {
	log := logger.SetLogger()
	cmd := exec.Command("sudo", "-S", "hostnamectl", "set-hostname", newName)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Cant get stdin:", err)
		return err
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println("Cant start comm:", err)
		return err

	}
	_, err = stdin.Write([]byte(password + "\n"))
	if err != nil {
		fmt.Println("Cant write:", err)
		return err
	}
	// Закрытие stdin, чтобы сигнализировать об окончании ввода
	err = stdin.Close()
	if err != nil {
		fmt.Println("Cant close stdin:", err)
		return err
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command waiting  error:", err)
		return err
	}

	cmd2 := exec.Command("hostname")
	out, err := cmd2.Output()
	if err != nil {
		log.Info("Access denied. Check Password ")
		return err
	}
	ok <- "Hostname changed successfully. New hostname is " + string(out)
	return nil
}
