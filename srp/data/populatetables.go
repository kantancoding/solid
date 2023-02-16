package main

import (
	"log"
	"math/rand"

	"srp/internal/db"
)

func main() {
	bank_id := 1
	db, err := db.Create("root", "")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for n := 1; n < 100; n++ {
		stmt, err := db.Prepare("INSERT INTO debit_transactions (amount, bank_id, process_id) VALUES (?,?,?);")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(rand.Intn(5000), bank_id, n)
		if err != nil {
			log.Fatal(err)
		}

		stmt, err = db.Prepare("INSERT INTO payment_transactions (amount, bank_id, process_id) VALUES (?,?,?);")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(rand.Intn(5000), bank_id, n)
		if err != nil {
			log.Fatal(err)
		}
	}
}
