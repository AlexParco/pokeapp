package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alexparco/pokeapp-api/config"
	_ "github.com/lib/pq"
)

type SqlClient struct {
	*sql.DB
}

func NewSqlClient(cfg *config.PostgresConfig) *SqlClient {
	connStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)

	db, err := sql.Open("postgres", connStr)
	fmt.Println(err)
	fmt.Println(db)

	if err != nil {
		log.Printf("Database Connection: %v", err)
		panic(err)
	}

	return &SqlClient{db}
}
