package main

import (
	"fmt"
	"sync"
)

type account struct {
	balance int
}

func (a *account) deposit(amount int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	a.balance += amount
	m.Unlock()
}

func (a *account) withdraw(amount int, m *sync.Mutex, wg *sync.WaitGroup) {
	m.Lock()
	defer wg.Done()
	defer m.Unlock()

	if amount > a.balance {
		fmt.Println("Not enough balance")
		return
	}
	a.balance -= amount
}

func main() {
	var m1 sync.Mutex
	var wg sync.WaitGroup

	a1 := account{500}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go a1.deposit(50, &m1, &wg)
	}


	for i := 0; i < 10; i++ {
		wg.Add(1)
		go a1.withdraw(100, &m1, &wg)
	}

	wg.Wait()
	fmt.Println(a1.balance)
}