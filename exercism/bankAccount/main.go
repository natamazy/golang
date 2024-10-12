package main

import (
	"math"
	"sync"
)

type Account struct {
	mu       *sync.Mutex
	balance  int64
	isClosed bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{
		balance:  amount,
		isClosed: false,
		mu:       &sync.Mutex{},
	}
}

func (a *Account) Balance() (int64, bool) {
	if a.isClosed {
		return 0, false
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.isClosed || (amount < 0 && a.balance < int64(math.Abs(float64(amount)))) {
		return 0, false
	}

	a.balance += amount

	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.isClosed {
		return 0, false
	}

	a.isClosed = true

	returnBalance := a.balance
	a.balance = 0
	return returnBalance, true
}
