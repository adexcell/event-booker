package repository

import "github.com/wb-go/wbf/dbpg"

type postgres struct {
	db *dbpg.DB
}

func NewPostgres(db *dbpg.DB) *postgres {
	return &postgres{db: db}
}
