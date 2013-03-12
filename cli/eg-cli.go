package main

import (
  "github.com/drio/drio.go/common/files"
  "github.com/drio/eg.go"
  "log"
  "net"
  "os"
)

func main() {
  probes := eg.Probes{}

  if len(os.Args) != 2 {
    panic("Usage: tool <probe_file_name>")
  }
  probes_fd, probes_rd := files.Xopen(os.Args[1])
  defer probes_fd.Close()
  probes.Init(probes_rd)

  service := ":8000"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  eg.CheckError(err)

  log.Printf("Starting server on port %s", service)
  listener, err := net.ListenTCP("tcp", tcpAddr)
  eg.CheckError(err)

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
