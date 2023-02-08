package reportfactory

import (
	"database/sql"
	"fmt"
)

type Report struct {
	BankID int
	Type   string
	Total  float64
	db     *sql.DB
}

func (r *Report) SaveReport() error {
	stmt, err := r.db.Prepare("INSERT INTO accounting_reports (bank_id, type, total) VALUES (?,?,?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(r.BankID, r.Type, r.Total)
	if err != nil {
		return err
	}

	return nil
}

func (r *Report) SendReport() error {
	fmt.Println("not implemented")
	return nil
}
