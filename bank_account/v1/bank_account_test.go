package v1

import (
	"sync"
	"testing"
)

type EventError struct {
	ET  Event
	err error
}

func checkError(t *testing.T, got error, expected *EventError, eventType Event) {
	t.Helper()

	if expected != nil && expected.ET == eventType {
		if got != expected.err {
			t.Errorf("Expected %v, got %v", expected.err, got)
		}
	} else {
		if got != nil {
			t.Errorf("Unexpected error: %v", got)
		}
	}
}

func TestBankAccount(t *testing.T) {
	depositWithdrawTestCases := []struct {
		desc            string
		deposits        []int
		withdrawals     []int
		overdraft       int
		eventError      *EventError
		expectedBalance int
	}{
		{"0 case", []int{}, []int{}, 0, nil, 0},
		{"unauthorized overdraft", []int{}, []int{10}, 0, &EventError{Withdrawal, ExceedingOverdraftError}, 0},
		{"authorized negative balance", []int{}, []int{10}, -50, nil, -10},
		{"authorized positive overdraft", []int{}, []int{}, 50, &EventError{UpdateOverdraft, ForbiddenPositiveOverdraftError}, 0},
		{"one deposit", []int{100}, []int{}, 0, nil, 100},
		{"negative deposit", []int{-50}, []int{}, 0, &EventError{Deposit, ForbiddenNegativeDepositError}, 0},
		{"negative withdrawal", []int{100}, []int{-10}, 0, &EventError{Withdrawal, ForbiddenNegativeWithdrawalError}, 100},
		{"one deposit + one withdrawal", []int{100}, []int{50}, 0, nil, 50},
		{"multiple withdrawals", []int{100}, []int{10, 10, 10, 10, 10}, 0, nil, 50},
		{"multiple deposits", []int{10, 10, 10, 10, 10}, []int{}, 0, nil, 50},
	}
	for _, tc := range depositWithdrawTestCases {
		t.Run(tc.desc, func(t *testing.T) {
			var wg sync.WaitGroup
			acc := NewAccount()
			err := acc.UpdateOverdraft(tc.overdraft)

			checkError(t, err, tc.eventError, UpdateOverdraft)

			for _, d := range tc.deposits {
				wg.Add(1)
				go func(d int) {
					defer wg.Done()
					err := acc.Deposit(d)

					checkError(t, err, tc.eventError, Deposit)
				}(d)
			}

			wg.Wait()

			for _, w := range tc.withdrawals {
				wg.Add(1)
				go func(w int) {
					defer wg.Done()
					err := acc.Withdraw(w)
					checkError(t, err, tc.eventError, Withdrawal)
				}(w)
			}

			wg.Wait()

			balance := acc.Balance()

			if balance != tc.expectedBalance {
				t.Errorf("Expected balance to be %d, got %d", tc.expectedBalance, balance)
			}

		})
	}

	t.Run("Transfers", func(t *testing.T) {
		account1 := NewAccount()
		account2 := NewAccount()
		account3 := NewAccount()

		account1.UpdateOverdraft(-100)
		account2.UpdateOverdraft(-100)

		account1.Deposit(100)
		account2.Deposit(100)
		account3.Deposit(20)

		type TestTransaction struct {
			from   *Account
			to     *Account
			amount int
			err    error
		}

		testCases := []struct {
			desc                string
			transactions        []TestTransaction
			expectedBalanceAcc1 int
			expectedBalanceAcc2 int
			expectedBalanceAcc3 int
		}{
			{"transfer $20 from 1 to 2", []TestTransaction{{account1, account2, 20, nil}}, 80, 100, 20},
			{"transfer $50 from 3 to 2", []TestTransaction{{account3, account2, 50, ExceedingOverdraftError}}, 100, 100, 20},
			{"transfer $50 from 1 to 1", []TestTransaction{{account1, account1, 50, ForbiddenSameAccountTransferError}}, 100, 100, 20},
		}

		for _, tc := range testCases {
			t.Run(tc.desc, func(t *testing.T) {
				var wg sync.WaitGroup

				for _, tx := range tc.transactions {
					wg.Add(1)
					go func(tx TestTransaction) {
						defer wg.Done()
						err := tx.from.Transfer(tx.to, tx.amount)

						if tx.err != err {
							t.Errorf("Expected %v, got %v", tx.err, err)
						}
					}(tx)
				}

				wg.Wait()
			})
		}

	})

}
