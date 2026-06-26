package model

type User struct {
	ID      int     `json:"user_id"`
	Name    string  `json:"user_name"`
	Balance float64 `json:"user_balance"`
}

func NewUser(name string, balance float64) User {
	return User{
		Name:    name,
		Balance: balance,
	}
}
