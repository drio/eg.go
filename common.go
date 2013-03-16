package eg

import (
  "fmt"
  "os"
)

func CheckError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}

// ReverseComplement performs a reverse complement on a
func ReverseComplement(ds *string) *string {
  db := []byte(*ds)
  rc := []byte{}
  for i := len(db) - 1; i >= 0; i-- {
    switch db[i] {
    case 'A':
      rc = append(rc, 'T')
    case 'C':
      rc = append(rc, 'G')
    case 'G':
      rc = append(rc, 'C')
    case 'T':
      rc = append(rc, 'A')
    case 'N':
      rc = append(rc, 'N')
    }
  }
  rcString := string(rc)
  return &rcString
}
