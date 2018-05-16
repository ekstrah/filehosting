package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0
	p := make([]byte, 2048)
	for {
		if i == 0 {
			msg := getUserAccount()
			i++
			buf := []byte(msg)
			_, err := Conn.Write(buf)
			if err != nil {
				fmt.Println(msg, err)
			}
			time.Sleep(time.Second * 1)
		} else {
			option := getUserOption()
			fmt.Println(option)
			buf := []byte(option)
			_, err := Conn.Write(buf)
			if err != nil {
				fmt.Println("NOOO", err)
			}
		}
		_, err = bufio.NewReader(Conn).Read(p)
		mesg := string(p[0:12])
		if err == nil {
			if mesg != "Successfully" {
				fmt.Printf("%s\n", mesg)
			}
		} else {
			fmt.Printf("Some error %v\n", err)
		}
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func getUserAccount() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Screen Name: ")
	text, _ := reader.ReadString('\n')
	return text
}

func getUserOption() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("1. Search")
	fmt.Println("2. One-to-One Direct Send")
	fmt.Println("3. Join Chat")
	fmt.Println("4. Help")
	fmt.Print("Option:  ")
	text, _ := reader.ReadString('\n')
	return text
}
