package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Account struct {
	sync.Mutex
	balance int
}

func (a *Account) AddMoney(money int) {
	a.Lock() // if discard mutex and run program with -trace, in the end we can finally see data race (for 3 tasks)
	a.balance += money
	fmt.Printf("Balance is %v.\n", a.balance)
	a.Unlock()
}

func main() {
	var (
		account = &Account{}
		wg      sync.WaitGroup
	)
	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			money := rand.Intn(100)
			fmt.Println(money)
			account.AddMoney(money)
		}()
	}
	wg.Wait()
}
