package tools

import (
	"bufio"
	"fmt"
	"github.com/GrosbergKirr/Server_hostname/internal/logger"
	"os"
	"os/exec"
	"strings"
)

func readCurrentDNSServers(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var dnsServers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "nameserver") {
			fields := strings.Fields(line)
			if len(fields) == 2 {
				dnsServers = append(dnsServers, fields[1])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %v", err)
	}
	return dnsServers, nil
}

func writeToFile(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}

func SetDNSServers(dnsServers string, password string, ok chan string) error {
	log := logger.SetLogger()
	const filepath = "/etc/resolv.conf"
	l := strings.Split(dnsServers, "/")

	currentDNSServers, err := readCurrentDNSServers(filepath)
	if err != nil {
		log.Info("Failed to read current DNS servers")
		return err
	}

	dnsMap := make(map[string]struct{})
	for _, dns := range currentDNSServers {
		dnsMap[dns] = struct{}{}
	}
	for _, dns := range l {
		dnsMap[dns] = struct{}{}
	}

	var d strings.Builder
	for dns := range dnsMap {
		d.WriteString("nameserver " + dns + "\n")
	}

	// Write dns to temporary file
	tmpFile := "/tmp/resolv.conf"
	if err := writeToFile(tmpFile, d.String()); err != nil {
		log.Info("failed to write to file: %v", err)
		return err
	}

	cmd := exec.Command("sudo", "-S", "mv", tmpFile, filepath)
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

	ok <- "DNS servers successfully updated"
	return nil
}
