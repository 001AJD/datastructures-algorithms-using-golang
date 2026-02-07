package main

import "fmt"

type Wallet struct {
	Owner   string
	Balance float32
}

func (w *Wallet) Deposit(amount float32) {
	w.Balance += amount
	fmt.Printf("%s deposited $%0.2f in the account, Current Balance is $%0.2f\n", w.Owner, amount, w.Balance)
}

func (w *Wallet) Spend(amount float32) {
	if w.Balance < amount {
		fmt.Println("Insufficient balance!")
	}
	w.Balance -= amount
	fmt.Printf("%s spent $%0.2f, available balance is $%0.2f\n", w.Owner, amount, w.Balance)
}

func (w *Wallet) GetCurrentBalance() float32 {
	return w.Balance
}
