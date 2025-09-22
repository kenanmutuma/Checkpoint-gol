package bank

var Balance = 1000

func Credit(amount int) {
	Balance += amount
}

func Debit(amount int) bool {
	if amount > Balance {
		return false
	}
	Balance -= amount
	return true
}
