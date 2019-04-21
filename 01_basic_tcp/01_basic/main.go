package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		fmt.Fprintf(conn, "hello...\n")
		fmt.Fprintln(conn, "how are you?")
		conn.Close()

	}

}
