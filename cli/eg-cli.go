package main

import (
  "github.com/drio/eg.go"
  "log"
  "net"
)

func main() {
  service := ":8000"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  eg.CheckError(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  eg.CheckError(err)

  log.Printf("Starting server on port %s", service)
  n := 0 // Number of connections since we started
  for {
    conn, err := listener.Accept()
    log.Printf("A new connection: %d\n", n)
    if err != nil {
      continue
    }
    go eg.HandleClient(conn, n)
    n++
  }
}
