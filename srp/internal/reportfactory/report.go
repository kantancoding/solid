package reportfactory

import (
	"database/sql"
	"fmt"
	"net/smtp"
)

type Report struct {
	BankID int
	Type   string
	Total  float64
	db     *sql.DB
}

func (r *Report) SaveReport() error {
	stmt, err := r.db.Prepare("INSERT INTO accounting_reports (bank_id, report_type, total) VALUES (?,?,?);")
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
	from := "accounting-report-service@gmail.com"
	password := "<Email Password>"

	to := []string{
		"accounting@company-name.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf("Bank ID: %d, Total: %g, Type: %s", r.BankID, r.Total, r.Type))
	fmt.Printf(fmt.Sprintf("Bank ID: %d, Total: %g, Type: %s", r.BankID, r.Total, r.Type))
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	fmt.Println("Email Sent Successfully!")

	return nil
}
