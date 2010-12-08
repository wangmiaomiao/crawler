package print

import (
  "os"
)

func PrintOrDie(msg string) {
  if _, err := os.Stdout.WriteString(msg + "\n"); err != nil {
    panic("Can't print msg: " + msg)
  }
}
