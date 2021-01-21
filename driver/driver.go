package driver

type DIALECTOR string

const (
	MySQL    = DIALECTOR("mysql")
	SQLite   = DIALECTOR("sqlite")
	Postgres = DIALECTOR("postgres")
	MSSQL    = DIALECTOR("mssql")
)
