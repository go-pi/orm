package driver

// DIALECTOR support mysql/sqlite/postgres/mssql
type DIALECTOR uint

const (
	MySQL    DIALECTOR = 1
	SQLite   DIALECTOR = 2
	Postgres DIALECTOR = 3
	MSSQL    DIALECTOR = 4
)

func (d DIALECTOR) String() string {
	switch d {
	case MySQL:
		return "mysql"
	case SQLite:
		return "sqlite"
	case Postgres:
		return "postgres"
	case MSSQL:
		return "mssql"
	}
	return "unknown"
}
