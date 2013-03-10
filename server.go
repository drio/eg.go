package eg

import (
	"net"
	"os"
	"log"
	"strconv"
)

func HandleClient(conn net.Conn, n int) {
  f, err := os.OpenFile(strconv.Itoa(n), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		log.Panicln("Problems opening file.")
		return
	}

	// close connection and file on exit
	defer f.Close()
	defer conn.Close()

	var buf [65536]byte
	for {
		// read upto 512 bytes
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		// write the n bytes read to file
		_, err = f.Write(buf[0:n])
		if err != nil {
			log.Panicln("Problems writing to file.")
			return
		}
	}
}

