package driver

// DIALECTOR support mysql/sqlite/postgres/mssql
type DIALECTOR string

const (
	MySQL    = DIALECTOR("mysql")
	SQLite   = DIALECTOR("sqlite")
	Postgres = DIALECTOR("postgres")
	MSSQL    = DIALECTOR("mssql")
)
