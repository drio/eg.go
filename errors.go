package eg

import (
  "fmt"
)

type egError struct {
  Where string
  What  string
}

func (e egError) Error() string {
  return fmt.Sprintf("Error in %s: %s ", e.Where, e.What)
}
