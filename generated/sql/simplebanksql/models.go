// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package simplebanksql

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Currencies string

const (
	CurrenciesUSD Currencies = "USD"
	CurrenciesEUR Currencies = "EUR"
	CurrenciesMXM Currencies = "MXM"
	CurrenciesCAD Currencies = "CAD"
	CurrenciesJPY Currencies = "JPY"
)

func (e *Currencies) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Currencies(s)
	case string:
		*e = Currencies(s)
	default:
		return fmt.Errorf("unsupported scan type for Currencies: %T", src)
	}
	return nil
}

type NullCurrencies struct {
	Currencies Currencies
	Valid      bool // Valid is true if Currencies is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCurrencies) Scan(value interface{}) error {
	if value == nil {
		ns.Currencies, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Currencies.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCurrencies) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Currencies), nil
}

type Account struct {
	ID         int64     `json:"id"`
	Owner      string    `json:"owner"`
	Balance    int64     `json:"balance"`
	CurrencyID int64     `json:"currency_id"`
	CreateadAt time.Time `json:"createad_at"`
}

type Currency struct {
	ID   int64      `json:"id"`
	Name Currencies `json:"name"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount     int64     `json:"amount"`
	CreateadAt time.Time `json:"createad_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreateadAt   time.Time `json:"createad_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount     int64     `json:"amount"`
	CreateadAt time.Time `json:"createad_at"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreateadAt        time.Time `json:"createad_at"`
}
