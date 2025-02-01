package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
  li, err := net.Listen("tcp", ":8080")
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
  scanner := bufio.NewScanner(conn)
  for scanner.Scan() {
    ln := scanner.Text()
    fmt.Println(ln)
  }

  defer conn.Close()
}