package main

import (
	"bank/db"
	"bank/pkg/core/services"
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"log"
	"os"
)

func main() {
///	database, err := sql.Open("sqlite3", "test")
	database, err := pgxpool.Connect(context.Background(), `postgres://dsurush:dsurush@localhost:5432/test?sslmode=disable`)
	if err != nil {
		log.Printf("Owibka - %e", err)
		log.Fatal("Can't Connection to DB")
	} else {
		fmt.Println("CONNECTION TO DB IS SUCCESS")
	}
	db.DbInit(database)
	Start(database)
}
const AuthorizationOperation = `1.Авторизация
2.Выйти`

func Start(database *pgxpool.Pool)  {
	for{
		fmt.Println(AuthorizationOperation)
		fmt.Println(`Выберите команду:`)
		var cmd int64
		fmt.Scan(&cmd)
		switch cmd {
		case 1:
			ok, id := services.Login(database)
			fmt.Println(ok)
			if ok {
				services.Authorized(database, id)
			}
		case 2:
			os.Exit(0)
		}
	}
}