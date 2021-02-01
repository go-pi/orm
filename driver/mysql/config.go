package mysql

import "fmt"

//Config mysql DNS params
type Config struct {
	HostName string
	HostPort string
	DBName   string
	UserName string
	Password string
	Charset  string
}

func (config Config) toDNS() string {
	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.UserName,
		config.Password,
		config.HostName,
		config.HostPort,
		config.DBName,
		config.Charset,
	)
	return dns
}
