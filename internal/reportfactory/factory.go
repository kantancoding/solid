package reportfactory

import (
	"database/sql"
	"errors"
)

type ReportFactory struct {
	db *sql.DB
}

func Construct(db *sql.DB) *ReportFactory {
	return &ReportFactory{db: db}
}

func (rf ReportFactory) Create(bankID int, reportType string) (*Report, error) {
	var total float64
	if reportType == "PAYMENT" {
		// aggregate all payment amounts for the given bank_id
		stmt, err := rf.db.Prepare("SELECT SUM(amount) FROM payment_transactions WHERE bank_id=?;")
		if err != nil {
			return nil, err
		}

		err = stmt.QueryRow(bankID).Scan(&total)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, sql.ErrNoRows
			}
			return nil, err
		}

		return &Report{BankID: bankID, Type: "PAYMENT", Total: total, db: rf.db}, nil
	} else if reportType == "DEBIT" {
		// aggregate all debit amounts for the given bank_id
		stmt, err := rf.db.Prepare("SELECT SUM(amount) FROM debit_transactions WHERE bank_id=?;")
		if err != nil {
			return nil, err
		}

		err = stmt.QueryRow(bankID).Scan(&total)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, sql.ErrNoRows
			}
			return nil, err
		}

		return &Report{BankID: bankID, Type: "DEBIT", Total: total, db: rf.db}, nil
	} else {
		return nil, errors.New("unsupported report type")
	}
}
