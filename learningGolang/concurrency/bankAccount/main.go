package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	mutex   *sync.RWMutex
	balance int
}

func (ba *BankAccount) deposit(amountToDep int) {
	ba.mutex.Lock()
	defer ba.mutex.Unlock()
	ba.balance += amountToDep
	fmt.Println("Deposited:\t\t", amountToDep, "\t| New Balance:", ba.balance)
}

func (ba *BankAccount) withdraw(amountToWithdraw int) {
	ba.mutex.Lock()
	defer ba.mutex.Unlock()

	if ba.balance >= amountToWithdraw {
		ba.balance -= amountToWithdraw
		fmt.Println("Withdrew:\t\t", amountToWithdraw, "\t| New Balance:", ba.balance)
		return
	}
	fmt.Println("Failed to withdraw:\t\t", amountToWithdraw, "\t| Balance:", ba.balance)
}

func main() {
	ba := &BankAccount{mutex: &sync.RWMutex{}}
	var wg sync.WaitGroup

	// Number of goroutines for deposits and withdrawals
	numRoutines := 10

	// Launch multiple goroutines to deposit
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ba.deposit(100)
		}()
	}

	// Launch multiple goroutines to withdraw
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ba.withdraw(50)
		}()
	}

	wg.Wait()

	// Final balance check
	fmt.Println("Final Balance:", ba.balance)
}
