package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		nett, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Path: ")
		path, _ := reader.ReadString('\n')
		nett.Write([]byte(path + "\n"))
		filename := nett.RemoteAddr().Network() + ".txt"
		fmt.Println(filename)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		n, _ := io.Copy(file, nett)
		nett.Close()
		fmt.Println(n)
		file.Close()

	}

}
