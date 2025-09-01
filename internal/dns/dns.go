package dns

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// SetGoogle changes DNS to Google DNS
func SetGoogle() error {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Detected OS: Windows")
		iface, err := getWindowsInterface()
		if err != nil {
			fmt.Println("Error detecting interface:", err)
			return err
		}
		cmd := exec.Command("netsh", "interface", "ipv4", "set", "dns", "name="+iface, "static", "8.8.8.8")
		if err := cmd.Run(); err != nil {
			fmt.Println("Error setting DNS:", err)
			return err
		}
		cmd = exec.Command("netsh", "interface", "ipv4", "add", "dns", "name="+iface, "8.8.4.4", "index=2")
		return cmd.Run()

	case "linux":
		fmt.Println("Detected OS: Linux")
		// Fallback: directly overwrite resolv.conf
		dnsConfig := "nameserver 8.8.8.8\nnameserver 8.8.4.4\n"
		return os.WriteFile("/etc/resolv.conf", []byte(dnsConfig), 0644)

	default:
		return fmt.Errorf("OS %s is not supported", runtime.GOOS)
	}
}

// SetDefault resets DNS back to default
func SetDefault() error {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Detected OS: Windows")
		iface, err := getWindowsInterface()
		if err != nil {
			return err
		}
		cmd := exec.Command("netsh", "interface", "ipv4", "set", "dns", "name="+iface, "dhcp")
		return cmd.Run()

	case "linux":
		fmt.Println("Detected OS: Linux")
		dnsConfig := "nameserver 127.0.0.53\n"
		return os.WriteFile("/etc/resolv.conf", []byte(dnsConfig), 0644)

	default:
		return fmt.Errorf("OS %s is not supported", runtime.GOOS)
	}
}

// helper: get active network interface on Windows
func getWindowsInterface() (string, error) {
	out, err := exec.Command("netsh", "interface", "show", "interface").Output()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(out), "\n")

	for _, line := range lines {
		if strings.Contains(line, "Connected") {
			fields := strings.Fields(line)
			if len(fields) > 3 {
				return fields[len(fields)-1], nil
			}
		}
	}

	// fallback if nothing found
	if strings.Contains(string(out), "Wi-Fi") {
		return "Wi-Fi", nil
	}
	if strings.Contains(string(out), "Ethernet") {
		return "Ethernet", nil
	}

	return "", fmt.Errorf("could not detect a valid network interface")
}

// GetCurrentDNS returns the current IPv4 DNS servers as a string
func GetCurrentDNS() (string, error) {
	fmt.Println("Fetching current DNS settings...")
	switch runtime.GOOS {
	case "windows":
		out, err := exec.Command("nslookup").Output()
		if err != nil {
			return "", err
		}
		return parseWindowsDNS(string(out)), nil

	case "linux":
		data, err := os.ReadFile("/etc/resolv.conf")
		if err != nil {
			return "", err
		}

		var dnsList []string
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "nameserver") {
				fields := strings.Fields(line)
				if len(fields) > 1 {
					dnsList = append(dnsList, fields[1])
				}
			}
		}
		if len(dnsList) == 0 {
			return "", fmt.Errorf("no DNS servers found in /etc/resolv.conf")
		}
		return strings.Join(dnsList, ", "), nil

	default:
		return "", fmt.Errorf("OS %s is not supported", runtime.GOOS)
	}
}

// Helper: parse nslookup output (Windows)
func parseWindowsDNS(output string) string {
	lines := strings.Split(output, "\n")
	var dns []string
	for _, l := range lines {
		if strings.Contains(l, "Address:") {
			parts := strings.Split(l, ":")
			if len(parts) > 1 {
				dns = append(dns, strings.TrimSpace(parts[1]))
			}
		}
	}
	if len(dns) == 0 {
		return "Unknown"
	}
	return strings.Join(dns, ", ")
}
