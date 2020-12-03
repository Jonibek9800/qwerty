package models

import (
	"bank/db"
	"context"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

type ATMs struct {
	ID int64
	Name string
	Status bool
}

func AddATM(database *pgxpool.Pool, address string) (ok bool, err error){
	_, err = database.Exec(context.Background(), db.AddNewATM, address)
	if err != nil {
		log.Println(`Can't insert to ATMs table new address, err is `, err)
		return false, err
	}
	return true, nil
}