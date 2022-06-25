package main

var deposits = make(chan int)
var balances = make(chan int)
var withdrawals = make(chan int)

func Deposit(amount int)  { deposits <- amount }
func Balance() int        { return <-balances }
func Withdraw(amount int) { withdrawals <- amount }

func teller() {
	var balance int

	for {
		select {
		case amount := <-withdrawals:
			if amount <= balance {
				balance -= amount
			}
		case amount := <-deposits:
			balance += amount

		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {

}
