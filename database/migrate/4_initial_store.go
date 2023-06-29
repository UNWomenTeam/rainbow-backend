package migrate

import (
	"fmt"

	"github.com/go-pg/migrations"
)

const addressTable = `
CREATE TABLE accounts (
id serial NOT NULL,
address_full text NOT NULL, 
PRIMARY KEY (id)
)`

const shopsTable = `
CREATE TABLE tokens (
id serial NOT NULL,
created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
account_id int NOT NULL REFERENCES accounts(id),
token text NOT NULL UNIQUE,
expiry timestamp with time zone NOT NULL,
mobile boolean NOT NULL DEFAULT FALSE,
identifier text,
PRIMARY KEY (id)
)`

func init() {
	up := []string{
		accountTable,
		tokenTable,
	}

	down := []string{
		`DROP TABLE tokens`,
		`DROP TABLE accounts`,
	}

	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating initial tables")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping initial tables")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
