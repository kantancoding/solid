package database

import (
	"database/sql"
)

type Database struct{}

func Construct() *Database {
	return &Database{}
}

func (database Database) PaymentAmountSum(bankID int, db *sql.DB) (float64, error) {
	var total float64
	// aggregate all payment amounts for the given bank_id
	stmt, err := db.Prepare("SELECT SUM(amount) FROM payment_transactions WHERE bank_id=?;")
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRow(bankID).Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, sql.ErrNoRows
		}
		return 0, err
	}

	return total, nil
}

func (database Database) DebitAmountSum(bankID int, db *sql.DB) (float64, error) {
	var total float64
	// aggregate all debit amounts for the given bank_id
	stmt, err := db.Prepare("SELECT SUM(amount) FROM debit_transactions WHERE bank_id=?;")
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRow(bankID).Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, sql.ErrNoRows
		}
		return 0, err
	}

	return total, nil
}
