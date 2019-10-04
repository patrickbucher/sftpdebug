package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	hostnameEnv = "SFTP_HOSTNAME"
	portEnv     = "SFTP_PORT"
	usernameEnv = "SFTP_USERNAME"
	passwordEnv = "SFTP_PASSWORD"
)

func main() {
	hostname := os.Getenv(hostnameEnv)
	port, err := strconv.Atoi(os.Getenv(portEnv))
	username := os.Getenv(usernameEnv)
	password := os.Getenv(passwordEnv)
	if hostname == "" || err != nil || username == "" || password == "" {
		fmt.Fprintln(os.Stderr, "missing connection parameters")
		os.Exit(1)
	}

	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey,
	}
	socket := fmt.Sprintf("%s:%d", hostname, port)
	sshConn, err := ssh.Dial("tcp", socket, sshConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "establish ssh connection: %w", err)
		os.Exit(1)
	}
	defer sshConn.Close()
	sftp, err := sftp.NewClient(sshConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "establish sftp connection: %w", err)
		os.Exit(1)
	}
	defer sftp.Close()
	fmt.Println(sftp)
}
