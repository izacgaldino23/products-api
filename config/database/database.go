package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/izacgaldino23/products-api/config"
	_ "github.com/lib/pq"
)

type Transaction struct {
	Builder *squirrel.StatementBuilderType
	Tx      *sql.Tx
}

func (t *Transaction) Commit() error {
	return t.Tx.Commit()
}

func (t *Transaction) Rollback() {
	_ = t.Tx.Rollback()
}

var DatabaseConnection *sql.DB

func OpenConnection() (err error) {
	connStr := fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", config.GetEnvironment().Database.Username, config.GetEnvironment().Database.Password, config.GetEnvironment().Database.Host, config.GetEnvironment().Database.Name)

	DatabaseConnection, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = DatabaseConnection.Ping()
	if err != nil {
		return err
	}

	return
}

func NewTransaction(readOnly bool) (transaction *Transaction, err error) {
	tx, err := DatabaseConnection.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: readOnly})
	if err != nil {
		return
	}

	builder := NewBuilder(tx)

	transaction = &Transaction{
		Builder: &builder,
		Tx:      tx,
	}

	return
}

func NewBuilder(tx *sql.Tx) squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		RunWith(tx)
}
