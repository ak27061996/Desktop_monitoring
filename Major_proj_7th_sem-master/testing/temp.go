package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const BUFFERSIZE = 1024

func main() {

	service := "127.0.0.1:7878"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	is_error(err)
	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	is_error(err)

	//message, _ := bufio.NewReader(conn).ReadString('\n')
	//fmt.Print("Message from server: " + message)

	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)15u0=-
	fmt.Println("successfully dialed connection")
	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")

	newFile, err := os.Create("./received/" + fileName)

	fmt.Println("creating newFile object")
	is_error(err)

	defer newFile.Close()
	var receivedBytes int64

	for {
		if (fileSize - receivedBytes) < BUFFERSIZE {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
			break
		}
		io.CopyN(newFile, connection, BUFFERSIZE)
		receivedBytes += BUFFERSIZE
	}
	fmt.Println("Received file completely!")
}

func is_error(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
