package tools

import (
	"bufio"
	"bytes"
	"fmt"
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

func SetDNSServers(dnsServers string, password string) error {
	const filepath = "/etc/resolv.conf"

	l := strings.Split(dnsServers, "/")
	fmt.Println(l)
	dnslist := strings.Join(l, "\n")
	fmt.Println(dnslist)

	currentDNSServers, err := readCurrentDNSServers(filepath)
	if err != nil {
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

	// Запись содержимого во временный файл
	tmpFile := "/tmp/resolv.conf"
	if err := writeToFile(tmpFile, d.String()); err != nil {
		return err
	}

	cmd := exec.Command("sudo", "-S", "mv", tmpFile, filepath)

	var stdin bytes.Buffer
	stdin.Write([]byte(password + "\n"))
	cmd.Stdin = &stdin

	var out strings.Builder
	var stderr strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to move resolv.conf: %v: %s", err, stderr.String())
	}

	fmt.Printf("Output: %s\n", out.String())
	return nil
}
