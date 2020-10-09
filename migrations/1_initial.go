package main

import (
	"fmt"
	"log"

	"github.com/go-pg/migrations/v7"
)

func init() {
	log.Printf("INIT MIGRATIONS")

	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table customer...")
		_, err := db.Exec(
			`CREATE TABLE customer(
					id bigserial CONSTRAINT pk PRIMARY KEY,
					first_name varchar(100),
					last_name varchar(100),
					patronymic_name varchar(100),
					phone varchar(20),
					email varchar(200)
					);
			insert into customer (first_name, last_name, patronymic_name, phone, email) values ('Клиент1', 'Клиентов1', 'Клиентович1', '77777777777', 'test1@test.ru');
			insert into customer (first_name, last_name, patronymic_name, phone, email) values ('Клиент2', 'Клиентов2', 'Клиентович2', '77777777777', 'test2@test.ru');
			insert into customer (first_name, last_name, patronymic_name, phone, email) values ('Клиент3', 'Клиентов3', 'Клиентович3', '77777777777', 'test3@test.ru');
			insert into customer (first_name, last_name, patronymic_name, phone, email) values ('Клиент4', 'Клиентов4', 'Клиентович4', '77777777777', 'test4@test.ru');
			insert into customer (first_name, last_name, patronymic_name, phone, email) values ('ДругойКлиент5', 'Клиентов5', 'Клиентович5', '77777777777', 'test5@test.ru');
`)
		log.Printf("CREATE customer")

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table customer...")
		_, err := db.Exec(`DROP TABLE customer;`)
		return err
	})
}
