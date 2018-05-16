package main

import (
	"fmt"
	"net"
	"os"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

/* Searching */
func messageSuccess(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("Successfully"), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

/* Searching */
func sendResult(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("1"), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

/* One-to-One Connection*/
func oneToOne(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("2"), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

/* Join Chat*/
func joinChat(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("3"), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		option := string(buf[0:n])
		fmt.Println("Received ", option, " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
		if option == "1" {
			go sendResult(ServerConn, addr)
		} else if option == "2" {
			go oneToOne(ServerConn, addr)
		} else if option == "3" {
			go joinChat(ServerConn, addr)
		}

		go messageSuccess(ServerConn, addr)
	}
}
