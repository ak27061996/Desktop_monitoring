package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
)

//constant buffer size
const BUFFERSIZE = 1024

//single function to rectify all errors
func is_error(err error) bool {
	if err != nil {
		fmt.Println("some error is there", err.Error)
		return true
	}
	return false
}

func main() {
	//server starts listening on port 27001

	//println(os.Args[1])

	//input address as command line argument (ex. 127.0.0.1:27001)
	listener, err := net.Listen("tcp", os.Args[1])
	is_error(err)

	//to close server automatically
	defer listener.Close()
	fmt.Println("###  Server started! Waiting for connections  ###")
	for {
		//extract connection from listening_port queue
		connection, err := listener.Accept()
		// is_error(err)
		if err != nil {
			fmt.Println("%#v", err)
		}
		//to automatically close client connection
		fmt.Println("Client connected")
		defer connection.Close()

		//to run shell script (schell script takes screenshot)
		// _, err = exec.Command("/home/anil/major_go_and_java/take_shot/take_shot.sh").Output()
		_, err = exec.Command("./take_shot/take_shot.sh").Output()
		is_error(err)

		//calling goroutines
		go serve_request(connection)
	}
}

func serve_request(connection net.Conn) {

	// fmt.Println("A client has connected!")

	//opening latest screenshot taken (every recent screenshot is named "snapshot.png")
	//and "./server_copies/" is the path relative to current working directory
	file, err := os.Open("./server_copies/snapshot.png")
	is_error(err)

	//Stat() returns information regarding size,name etc about file opened
	//file_info, err := file.Stat()
	// is_error(err)

	//file_size := fill_buffer(strconv.FormatInt(file_info.Size(), 10), 10)
	//file_name := fill_buffer(file_info.Name(), 64)

	// fmt.Println("Sending filename and filesize")

	// connection.Write([]byte(file_size))
	//connection.Write([]byte(file_name))
	// println("file_size: ", file_size)

	//creating buffer (into which file is read then sent onto "connection")
	sendBuffer := make([]byte, BUFFERSIZE)

	// fmt.Println("Start sending file!")
	for {
		//read content of image file into buffer
		_, err = file.Read(sendBuffer)
		//read  until END_OF_FILE
		if err == io.EOF {
			break
		}
		is_error(err)
		//write content into buffer
		connection.Write(sendBuffer)
	}
	fmt.Println(" file sent ")
	return
}

//function to fill remaining buffer with addtional characters
func fill_buffer(rec_string string, length int) string {
	for {
		lengtString := len(rec_string)
		if lengtString < length {
			rec_string = rec_string + ":"
			continue
		}
		break
	}
	return rec_string
}
