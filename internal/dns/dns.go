package dns

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Đổi DNS sang Google
func SetGoogle() error {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("netsh", "interface", "ip", "set", "dns", "name=Ethernet", "static", "8.8.8.8")
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("netsh", "interface", "ip", "add", "dns", "name=Ethernet", "8.8.4.4", "index=2")
		return cmd.Run()
	case "linux":
		dnsConfig := "nameserver 8.8.8.8\nnameserver 8.8.4.4\n"
		return os.WriteFile("/etc/resolv.conf", []byte(dnsConfig), 0644)
	default:
		return fmt.Errorf("OS %s chưa hỗ trợ", runtime.GOOS)
	}
}

// Đổi DNS về mặc định
func SetDefault() error {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("netsh", "interface", "ip", "set", "dns", "name=Ethernet", "dhcp")
		return cmd.Run()
	case "linux":
		dnsConfig := "nameserver 127.0.0.53\n"
		return os.WriteFile("/etc/resolv.conf", []byte(dnsConfig), 0644)
	default:
		return fmt.Errorf("OS %s chưa hỗ trợ", runtime.GOOS)
	}
}
