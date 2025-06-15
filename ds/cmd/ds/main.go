package main

import (
	"fmt"

	"ds/helper"
	"ds/pq"
	"ds/trymap"
	"ds/trymap/fuckup"

	"github.com/google/btree"
)

type IntItem int

func (a IntItem) Less(b btree.Item) bool {
	return a < b.(IntItem)
}

func main() {
	fmt.Printf("%#v\n", fuckup.Solution(helper.Qc("ADD 1", "EXISTS 1", "ADD 2", "ADD 2", "EXISTS 2", "REMOVE 2", "EXISTS 2", "ADD 3", "GET_NEXT 2")))
	trymap.TrySyncMap()
	trymap.TryTMap()

	tr := btree.New(2)
	for _, v := range []int{5, 1, 7, 3, 2, 4, 6} {
		tr.ReplaceOrInsert(IntItem(v))
	}

	if tr.Has(IntItem(3)) {
		fmt.Println("Found 3")
	} else {
		fmt.Println("Did not find 3")
	}

	tr.Ascend(func(i btree.Item) bool {
		fmt.Println(i)
		return true
	})

	println("Priority Queue Example:")
	pq.Example()
}
