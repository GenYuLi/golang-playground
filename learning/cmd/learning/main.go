package main

import (
  "fmt"
	"learning/bank"
)

func main() {
	list := []int64{10, 100, 20, 50, 30}
	bank := bank.Constructor(list)
	list[0] = 30
	bank.Withdraw(5, 20)
  
	fmt.Printf("%#v\n", bank)          // bank.Bank{account:[]int64{10, 100, -10, 50, 30}}
	fmt.Printf("%+v\n", bank)         // {account:[10 100 -10 50 30]}
	fmt.Println(bank)                  // {account:[10 100 -10 50 30]}
}
