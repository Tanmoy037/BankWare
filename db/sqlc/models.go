// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"time"
)

type Account struct {
	ID        int64     `db:"id"`
	Owner     string    `db:"owner"`
	Balance   int64     `db:"balance"`
	Currency  string    `db:"currency"`
	CreatedAt time.Time `db:"created_at"`
}

type Entry struct {
	ID        int64 `db:"id"`
	AccountID int64 `db:"account_id"`
	// can be negative or positive
	Amount    int64     `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}

type Transfer struct {
	ID            int64 `db:"id"`
	FromAccountID int64 `db:"from_account_id"`
	ToAccountID   int64 `db:"to_account_id"`
	// must be positive
	Amount    int64     `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}
