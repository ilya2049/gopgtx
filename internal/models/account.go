package models

import "fmt"

type Account struct {
	ID      int
	Balance float64
}

func (a Account) String() string {
	return fmt.Sprintf("[%d] %.2f", a.ID, a.Balance)
}
