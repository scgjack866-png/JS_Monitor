package net

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Ullaakut/nmap/v3"
	"log"
	"os/exec"
	"strings"
	"time"
)

var (
	Normal          = 100
	HostUnreachable = 101
	TCPFailed       = 201
	UDPFailed       = 202
	DomesticError   = 301
	ForeignError    = 302
)

func CheckIPTCPAndUDP(IP, port, password, method, localPort, tunnel, authName string) int {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5-minute timeout.
	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets(IP),
		nmap.WithPorts(port),
		nmap.WithUDPScan(),
		nmap.WithTCPNullScan(),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		log.Printf("run finished with warnings: %s\n", *warnings) // Warnings are non-critical errors from nmap.
	}
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			if port.Protocol == "tcp" {
				fmt.Println(port.Protocol)
				fmt.Println(port.State.State)
				if !strings.Contains(port.State.State, "open") {
					return TCPFailed

				}
			} else {
				fmt.Println(port.Protocol)
				fmt.Println(port.State.State)
				if !strings.Contains(port.State.State, "open") {
					return UDPFailed
				}
			}
		}
		return startServers(IP, port, password, method, localPort, tunnel, authName)
	}
	return HostUnreachable
}

func startServers(IP, port, password, method, localPort, tunnel, authName string) int {
	var innerUrl = "ip.sb"
	var outUrl = "ip.sb"
	fmt.Println("/usr/local/mynet/myss-linux-amd64 -sd \"" + tunnel + "\" -c myss://" + method + ":" + password + "@" + IP + ":" + port + " -socks 127.0.0.1:" + localPort + " -user " + authName + " >/dev/null 2>&1 &")
	cmd := exec.Command("bash", "-c", "/usr/local/mynet/myss-linux-amd64 -sd \""+tunnel+"\" -c myss://"+method+":"+password+"@"+IP+":"+port+" -socks 127.0.0.1:"+localPort+" -user "+authName+" >/dev/null 2>&1 &")
	time.Sleep(2 * time.Second)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println("curl --socks5 127.0.0.1:" + localPort + " -k -m 10 -s " + innerUrl)
	innerCmd := exec.Command("bash", "-c", "curl --socks5 127.0.0.1:"+localPort+" -k -m 10 -s "+innerUrl)
	var out1 bytes.Buffer
	var stderr bytes.Buffer
	innerCmd.Stdout = &out1
	innerCmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error(), stderr.String())
	}
	fmt.Printf("%s", out)

	outCmd := exec.Command("bash", "-c", "curl --socks5 127.0.0.1:"+localPort+" -k -m 10 -s "+outUrl)
	out, err = outCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s", out)
	return Normal
}
