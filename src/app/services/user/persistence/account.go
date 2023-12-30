package persistence

type Account struct {
	Login    string
	Password string
}

var Tessst *Account = &Account{Login: "123", Password: "awfw"}
