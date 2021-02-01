package mysql

import (
	"github.com/go-pi/orm/driver"
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	dns    string
	config Config
}

func (ms mysql) Name() string {
	return driver.MySQL.String()
}

func (ms mysql) DNS() string {
	return ms.dns
}

/*
// Config MySQL config information
type Config struct {
	DriverName                string
	DSN                       string
	SkipInitializeWithVersion bool
	DefaultStringSize         uint
	DefaultDatetimePrecision  *int
	DisableDatetimePrecision  bool
	DontSupportRenameIndex    bool
	DontSupportRenameColumn   bool
	DontSupportForShareClause bool
}
*/

func New(arg interface{}) driver.Driver {
	switch val := arg.(type) {
	case string:
		println("string", val)
		return &mysql{
			dns: arg.(string),
		}
	case Config:
		config := arg.(Config)
		return &mysql{
			dns:    config.toDNS(),
			config: config,
		}
	}
	println("is nil")
	return nil
}
