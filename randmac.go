package main

import (
	"crypto/rand"
	"fmt"
	"net"
	"os/exec"
	"runtime"
)

func randomMAC() (net.HardwareAddr, error) {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	// Set the local bit
	buf[0] |= 2
	return net.HardwareAddr(buf), nil
}

func main() {
	// Generate a random MAC address
	mac, err := randomMAC()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Changing MAC address to %s\n", mac)

	// Identify the OS
	switch os := runtime.GOOS; os {
	case "darwin":
		// MacOS
		cmd := exec.Command("sudo", "ifconfig", "en0", "ether", mac.String())
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	case "linux":
		// Linux
		cmd := exec.Command("sudo", "ifconfig", "eth0", "ether", mac.String())
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	case "windows":
		// Windows (uses PowerShell and Netsh, adjust interface name accordingly)
		cmd := exec.Command("powershell", "-Command", "Start-Process", "cmd", "/C", fmt.Sprintf(`netsh interface set interface "Ethernet" admin=disable && netsh interface set interface "Ethernet" admin=enable && netsh interface set interface "Ethernet" newmac=%s`, mac))
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	default:
		fmt.Printf("%s.\n", "Unsupported operating system")
	}
}
