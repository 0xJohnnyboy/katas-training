package v1

import (
	"errors"
	"sync"
)

var ExceedingOverdraftError = errors.New("Exceeding overdraft")
var ForbiddenNegativeDepositError = errors.New("Deposit must be over $0")
var ForbiddenNegativeWithdrawalError = errors.New("Withdrawal must be over $0")
var ForbiddenNegativeTransferError = errors.New("Transfer must be over $0")
var ForbiddenPositiveOverdraftError = errors.New("Overdraft must <= 0")
var ForbiddenSameAccountTransferError = errors.New("Can't transfer to the same account")

type Event int

const (
	Deposit Event = iota
	Withdrawal
	CheckBalance
	UpdateOverdraft
	Transfer
)

type Tx struct {
	Event  Event
	From   *Account
	To     *Account
	Amount int
}

type Account struct {
	mu sync.RWMutex
	// id           string
	transactions []Tx
	overdraft    int
}

func NewAccount() *Account {
	return &Account{
		// id:           "test",
		transactions: []Tx{},
		overdraft:    0,
	}
}

func (a *Account) UpdateOverdraft(amount int) error {
	if amount > 0 {
		return ForbiddenPositiveOverdraftError
	}
	a.mu.Lock()
	defer a.mu.Unlock()

	a.overdraft = amount
	tx := Tx{UpdateOverdraft, nil, nil, amount}
	a.transactions = append(a.transactions, tx)

	return nil
}

func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return ForbiddenNegativeDepositError
	}
	a.mu.Lock()
	defer a.mu.Unlock()

	tx := Tx{Deposit, nil, a, amount}
	a.transactions = append(a.transactions, tx)

	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return ForbiddenNegativeWithdrawalError
	}
	a.mu.Lock()
	defer a.mu.Unlock()

	balance := a.getBalance()
	if balance-amount < a.overdraft {
		return ExceedingOverdraftError
	}

	tx := Tx{Withdrawal, a, nil, amount}
	a.transactions = append(a.transactions, tx)

	return nil
}

func (a *Account) Transfer(to *Account, amount int) error {
	if a == to {
		return ForbiddenSameAccountTransferError
	}
	if amount <= 0 {
		return ForbiddenNegativeTransferError
	}

	a.mu.Lock()
	to.mu.Lock()
	defer a.mu.Unlock()
	defer to.mu.Unlock()

	balance := a.getBalance()

	if balance-amount < a.overdraft {
		return ExceedingOverdraftError
	}

	tx := Tx{Transfer, a, to, amount}

	a.transactions = append(a.transactions, tx)
	to.transactions = append(a.transactions, tx)

	return nil
}

func (a *Account) Balance() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	sum := 0

	for _, tx := range a.transactions {
		switch tx.Event {
		case Deposit:
			sum += tx.Amount
		case Withdrawal:
			sum -= tx.Amount
		case Transfer:
			if tx.From == a {
				sum -= tx.Amount
			} else {
				sum += tx.Amount
			}
		}
	}

	tx := Tx{CheckBalance, nil, nil, 0}
	a.transactions = append(a.transactions, tx)

	return sum
}

func (a *Account) getBalance() int {
	sum := 0

	for _, e := range a.transactions {
		switch e.Event {
		case Deposit:
			sum += e.Amount
		case Withdrawal:
			sum -= e.Amount
		}
	}

	return sum
}
