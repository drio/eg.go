package main

import (
  "fmt"
  "github.com/drio/drio.go/bio/fasta"
  "github.com/drio/drio.go/common/files"
  "github.com/drio/eg.go"
  "log"
  "net"
  "os"
)

func main() {
  if len(os.Args) != 3 {
    fmt.Fprintf(os.Stderr, "Usage: tool <probe_file_name> <reads.fa>\n")
    os.Exit(1)
  }

  log.Printf("Loading probes.")
  probes := loadProbes()
  log.Printf("Number of probes loaded: %d\n", probes.NumLoaded())

  // start server or screen
  //startServer()

  reads_fn := os.Args[2]
  fpReads, readsReader := files.Xopen(reads_fn)
  defer fpReads.Close()
  var fqr fasta.FqReader
  fqr.Reader = readsReader
  eg.Compute(fqr, probes, os.Stdin)
}

func loadProbes() eg.Probes {
  probes := eg.Probes{}
  fd, rd := files.Xopen(os.Args[1]) // filedesc, reader for probes
  probes.Load(rd)
  fd.Close()
  return probes
}

func startServer() {
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
