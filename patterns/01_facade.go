package main

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"fmt"
)

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
}

func NewWalletFacade() *WalletFacade {
	facade := &WalletFacade{
		account:      NewAccount("myAccountID123"),
		wallet:       NewWallet(),
		securityCode: NewSecurityCode("myHash32"),
	}
	return facade
}

func (f *WalletFacade) AddMoneyToWallet(accountID string, securityCode string, amount int) error {
	fmt.Println("Starting a transaction...")

	err := f.account.Verify(accountID)
	if err != nil {
		return fmt.Errorf("Account validation failed! Reason: %s", err)
	}

	err = f.securityCode.CheckSecurity(securityCode)
	if err != nil {
		return fmt.Errorf("Security validation failed! Reason: %s", err)
	}

	fmt.Println("AccountID: \t", accountID)
	fmt.Println("SecurityCode: \t", securityCode)
	fmt.Println("Money before: \t", f.wallet.balance)

	err = f.wallet.AddMoney(amount)
	if err != nil {
		return fmt.Errorf("Transaction failed! Reason: %s", err)
	}

	fmt.Println("Transaction completed!")
	fmt.Println("Money after: \t", f.wallet.balance)

	return nil
}

type Account struct {
	ID      string
	blocked bool
}

func NewAccount(id string) *Account {
	return &Account{
		ID:      id,
		blocked: false,
	}
}
func (a *Account) Verify(accountID string) error {
	if a.ID != accountID {
		return fmt.Errorf("AccountID is not correct!")
	}
	return nil
}

type Wallet struct {
	balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}
func (w *Wallet) AddMoney(value int) error {
	w.balance += value
	return nil
}
func (w *Wallet) SubstractMoney(value int) error {
	w.balance -= value
	return nil
}

type SecurityCode struct {
	hash string
}

func NewSecurityCode(newHash string) *SecurityCode {
	return &SecurityCode{
		hash: newHash,
	}
}
func (c *SecurityCode) CheckSecurity(code string) error {
	if c.hash != code {
		return fmt.Errorf("SecurityCode is not correct!")
	}
	return nil
}

func main() {
	facade := NewWalletFacade()

	err := facade.AddMoneyToWallet("myAccountID123", "myHash32", 1000)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()

	err = facade.AddMoneyToWallet("wrongID", "myHash32", 1000)
	if err != nil {
		fmt.Println(err)
	}
}
