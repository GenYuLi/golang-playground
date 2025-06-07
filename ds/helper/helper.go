package helper

import (
  "strings"
)

func Qc(lines ...string) [][]string {
  out := make([][]string, len(lines))
  for i,ln := range lines {
    out[i] = strings.Fields(ln)
  }
  return out
}
