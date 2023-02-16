package main

import (
	"log"
	"net/http"
	"strconv"

	"srp/internal/database"
	"srp/internal/db"
	"srp/internal/reportfactory"
)

func main() {
	db, err := db.Create("root", "")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	reportFactory := reportfactory.Construct(db)

	http.HandleFunc("/v1/report", func(w http.ResponseWriter, r *http.Request) {
		bankIDString := r.URL.Query().Get("bank_id")
		reportType := r.URL.Query().Get("report_type")

		if bankIDString == "" || reportType == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		bankID, err := strconv.Atoi(bankIDString)
		if err != nil {
			log.Print(err)
		}

		report, err := reportFactory.Create(bankID, reportType)
		if err != nil {
			log.Print(err)
		}

		database := database.Construct()
		err = database.SaveReport(report.BankID, report.Type, report.Total, db)
		if err != nil {
			log.Print(err)
		}

		err = report.SendReport()
		if err != nil {
			log.Print(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
