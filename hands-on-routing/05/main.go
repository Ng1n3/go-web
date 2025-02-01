


package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":3050")
	if err != nil {
		log.Fatal(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	io.WriteString(conn, "I see ou connected.\n")
	fmt.Fprintln(conn, "I see you connected.")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}

	defer conn.Close()
  io.WriteString(conn, "Writting to the response.")
  
  
  


body := "Writting to the response."

// Write the status line.
io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
// Write the Content-Length header.
fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
// Write the Content-Type header.
fmt.Fprint(conn, "Content-Type: text/plain\r\n")
// Write the blank line to indicate the end of the headers.
io.WriteString(conn, "\r\n")
// Write the response body.
io.WriteString(conn, body)
}