package orm

import (
	"context"
	"database/sql"
)

//Dialector ORM database dialector
type Dialector interface {
	Name()
	Initialize(db *DB) error
	Migrator(db *DB) Migrator
	DataTypeOf() string
	DafaultValueOf() string
	BindVarTo()
	QuoteTo()
	Explain() string
}

//Migrator ORM database migrator
type Migrator interface {
}

// ConnPool db conns pool interface
type ConnPool interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
