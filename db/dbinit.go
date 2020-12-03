package db

import (
	"context"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

func DbInit(database *pgxpool.Pool) {
	DDLs := []string{CreateUsersAccount, CreateATMsTable}
	for _, ddl := range DDLs{
		_, err := database.Exec(context.Background(), ddl)
		if err != nil {
			log.Fatalf("Can't init %s err is %e", ddl, err)
		}
	}
}
