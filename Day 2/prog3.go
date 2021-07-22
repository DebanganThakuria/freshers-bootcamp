package main

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	balance int
}

func (a *account) deposit(amount int, m *sync.Mutex) {
	m.Lock()
	a.balance += amount
	m.Unlock()
}

func (a *account) withdraw(amount int, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	if amount > a.balance {
		fmt.Println("Not enough balance")
		return
	}
	a.balance -= amount
}

func main() {
	var m1 sync.Mutex

	a1 := account{500}

	for i := 0; i < 10; i++ {
		go a1.deposit(50, &m1)
	}


	for i := 0; i < 10; i++ {
		go a1.withdraw(100, &m1)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(a1.balance)
}