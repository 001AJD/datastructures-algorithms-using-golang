package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Pointers lesson")
	var pointer *uint8
	var num uint8
	num = 25
	pointer = &num
	fmt.Println("Actual num variable", num)
	fmt.Println("pointer variable contains the address of num variable", pointer)
	fmt.Println("using * operator to access the value at the address stored in pointer", *pointer)

	var familyWallet *Wallet = &Wallet{
		Owner:   "Dhomne family",
		Balance: 500.00,
	}

	fmt.Printf("Initial Balance in the account :: $%0.2f\n", familyWallet.GetCurrentBalance())
	familyWallet.Deposit(200)
	familyWallet.Spend(300)
	fmt.Printf("Current Balance :: $%0.2f\n", familyWallet.GetCurrentBalance())

	var pointer2 *uint8 = new(uint8)
	*pointer2 = 25
	fmt.Println("Address of pointer2 ", pointer2)
	fmt.Println("Value at pointer2", *pointer2)

}
