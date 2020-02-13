package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rudyrobles"
	password = "your-password"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT *
		FROM users
		INNER JOIN orders ON users.id=orders.user_id`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var userID, orderID, amount int
		var email, name, desc string
		if err := rows.Scan(&userID, &name, &email, &orderID, &userID, &amount, &desc); err != nil {
			panic(err)
		}
		fmt.Println("userID:", userID, "name:", name, "email:", email, "orderID:", orderID, "amount:", amount, "desc:", desc)
	}
	if rows.Err() != nil {
		panic(rows.Err())
	}
}
