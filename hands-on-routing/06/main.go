package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
  	var i int
	var rMethod, rURI string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(ln)
      // fmt.Println(xs)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
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
