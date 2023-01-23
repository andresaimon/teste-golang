package db

import "database/sql"

type Store interface {
	Querier
}

// tipagem para a realização da conexão com o banco de dados e as requisições feitas:
type ExecuteStore struct {
	db *sql.DB
	*Queries
}

// definição da função a ser utilizada nos testes unitários e nas requisições HTTP
func ExecuteNewStore(db *sql.DB) *ExecuteStore {
	return &ExecuteStore{
		db:      db,
		Queries: New(db),
	}
}
