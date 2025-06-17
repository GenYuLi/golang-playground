package bank

import "slices"

type Bank struct {
	account []int64
}

func Constructor(balance []int64) Bank {
	return Bank{
		// reference to the slice, so if you might modify the balance slice outside this bank
		// account: balance, // this will not work as expected, because it will modify the original slice

		// a method to deep copy the slice, ... is a syntax to unpack the slice
		// account: append([]int64{}, balance...),

		account: slices.Clone(balance), // slices.copy is available in Go 1.19+
	}
}

func (b *Bank) Transfer(ac1 int, ac2 int, money int64) bool {
	if ac1 > len((*b).account) || ac2 >= len((*b).account) {
		return false
	}
	if money > (*b).account[ac1-1] {
		return false
	}
	(*b).account[ac1-1] -= money
	(*b).account[ac2-1] += money
	return true
}

func (b *Bank) Deposit(ac int, money int64) bool {
	if ac > len((*b).account) {
		return false
	}
	(*b).account[ac-1] += money
	return true
}

func (b *Bank) Withdraw(ac int, money int64) bool {
	if ac > len((*b).account) {
		return false
	}

	if money > (*b).account[ac-1] {
		println((*b).account[ac-1])
		return false
	}

	(*b).account[ac-1] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
