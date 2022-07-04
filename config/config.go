package config

type Account struct {
	IDType   string
	Username string
	Password string
}

type Config struct {
	Account *Account
}
