package main

import "fmt"

type Bank struct {
	balance       []int64
	accountLength int
}

func Constructor(balance []int64) Bank {
	len := len(balance)
	return Bank{balance, len}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > this.accountLength || account2 > this.accountLength || this.balance[account1-1]-money < 0 {
		return false
	}
	this.balance[account1-1] -= money
	this.balance[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > this.accountLength {
		return false
	}
	this.balance[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if this.accountLength < account || this.balance[account-1] < money {
		return false
	}
	this.balance[account-1] -= money
	return true
}

func main() {
	balance := []int64{10, 100, 20, 50, 30}

	obj := Constructor(balance)
	param_1 := obj.Withdraw(3, 10)
	param_2 := obj.Transfer(5, 1, 20)
	param_3 := obj.Deposit(5, 20)
	param_4 := obj.Transfer(3, 4, 15)
	param_5 := obj.Withdraw(10, 50)
	fmt.Println(param_1, param_2, param_3, param_4, param_5)
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
