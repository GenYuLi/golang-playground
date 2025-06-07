package main

import (
	"ds/trymap"
	"ds/trymap/fuckup"
	"fmt"
  "ds/helper"
)




func main() {

  fmt.Printf("%#v\n", fuckup.Solution(helper.Qc("ADD 1", "EXISTS 1", "ADD 2", "ADD 2", "EXISTS 2", "REMOVE 2", "EXISTS 2", "ADD 3", "GET_NEXT 2"))) 
  trymap.TrySyncMap()
  trymap.TryTMap()
}
