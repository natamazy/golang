package main

import (
	"fmt"
	"sync"
)

type inventory struct {
	mutex *sync.RWMutex
	goods map[string]uint
}

func (in *inventory) checkInventory() {
	in.mutex.RLock()
	defer in.mutex.RUnlock()

	for stockName, stockCount := range in.goods {
		fmt.Printf("%s: %d pieces\n", stockName, stockCount)
	}
}

func (in *inventory) addStock(stockName string, stockCountToAdd uint) {
	in.mutex.Lock()
	defer in.mutex.Unlock()

	if stockCountToAdd == 0 {
		fmt.Println("Stock count need to be more than 0")
		return
	}

	_, ok := in.goods[stockName]
	if !ok {
		in.goods[stockName] = 0
	} else {
		in.goods[stockName] += stockCountToAdd
	}
}

func (in *inventory) removeStock(stockName string, stockCountToRemove uint) {
	in.mutex.Lock()
	defer in.mutex.Unlock()

	if stockCountToRemove == 0 {
		fmt.Println("Stock count need to be more than 0")
		return
	}

	currecntCount, ok := in.goods[stockName]
	if !ok {
		fmt.Println("No such item:", stockName)
		return
	}

	if currecntCount < stockCountToRemove {
		fmt.Println("Current stock count is not enough")
		return
	}

	in.goods[stockName] -= currecntCount
}

func main() {
	inv := &inventory{
		mutex: &sync.RWMutex{},
		goods: make(map[string]uint),
	}

	// WaitGroup to ensure all goroutines complete
	var wg sync.WaitGroup

	// Number of goroutines
	numRoutines := 10

	// Add stock concurrently
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			inv.addStock("apple\t\t", uint(id+1))
			fmt.Println("Added stock in goroutine\t\t", id)
		}(i)
	}

	// Remove stock concurrently
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			inv.removeStock("apple\t\t", uint(id))
			fmt.Println("Removed stock in goroutine\t\t", id)
		}(i)
	}

	// Check inventory concurrently
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			inv.checkInventory()
			fmt.Println("Checked inventory in goroutine\t\t", id)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("Final Inventory:")
	inv.checkInventory()
}
