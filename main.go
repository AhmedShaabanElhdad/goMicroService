package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	config := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPasswd,
		Addr:                 Envs.DBAddr,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewSqlStorage(config)
	db, err := sqlStorage.init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)

	api := NewServer(":3000", store)
	api.Serve()
}
