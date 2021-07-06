package utils

import (
	"fmt"
	"net"
	"strconv"
)

func GetAvailablePort() (string, error) {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:0", "0.0.0.0"))
	if err != nil {
		return "", err
	}
	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		return "", err
	}
	defer listener.Close()
	return strconv.Itoa(listener.Addr().(*net.TCPAddr).Port), nil
}
