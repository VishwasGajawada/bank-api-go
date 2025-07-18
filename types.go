package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber int64     `json:"accountNumber"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"createdAt"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: int64(rand.Intn(10000000)),
		CreatedAt:     time.Now().UTC(),
	}
}
