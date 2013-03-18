package eg

import (
  "bufio"
  "github.com/drio/drio.go/bio/fasta"
  "log"
  "net"
  "os"
  "strconv"
)

func HandleClient(conn net.Conn, n int, probes Probes) {
  log.Printf("A new connection: %d", n)

  // Open the output file
  fdOut, err := os.OpenFile(strconv.Itoa(n), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
  if err != nil {
    log.Panicln("Problems opening file.")
    return
  }

  // close connection and file on exit
  defer fdOut.Close()
  defer conn.Close()

  // Create a buffer for the reader (conn - socket)
  bReader := bufio.NewReader(conn)
  // Now create a FastaQ reader
  var fqr fasta.FqReader
  fqr.Reader = bReader
  Compute(fqr, probes, fdOut)

  log.Printf("Done with connection %d.", n)
}
