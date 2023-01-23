package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// configuração do banco de dados e o usuário/senha
const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_products?sslmode=disable"
)

var testQueries *Queries

// função para testar a conexão com o banco de dados
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
