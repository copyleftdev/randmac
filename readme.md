

# randmac

A simple Go application to change the MAC address of your network interface. This app generates a random MAC address and attempts to set it as the MAC address for your network interface.

Supports Linux, MacOS, and Windows.

## Requirements

- Go installed on your machine. You can download Go [here](https://golang.org/dl/).
- Administrator or root permissions are required to change MAC addresses.
- Know the name of the network interface you wish to change the MAC address of. This name varies by operating system and network configuration.

## How to Run

1. Clone this repository to your local machine.
2. Open the `randmac.go` file and replace the network interface name in the script with your own network interface name.
3. From the terminal, navigate to the directory containing the `randmac.go` file.
4. Run the command `go run randmac.go`.

Please note that some systems require you to restart your network interface or even your computer for the new MAC address to take effect.

## Warning

Use this app with caution. Changing your MAC address can lead to network connectivity issues, and in some cases, it might even violate terms of service or laws. Always ensure you understand what you're doing when changing MAC addresses.


