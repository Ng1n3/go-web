package main

import (
	"io"
	"log"
	"net"
)

func main() {
  sv, err := net.Listen("tcp", ":3050")
  if err != nil {
    log.Panic(err)
  }

  defer sv.Close()

  for {
    conn, err := sv.Accept()
    if err != nil {
      log.Println(err)
      continue
    }

    io.WriteString(conn, "TCP network connected")

    conn.Close()
  }
}