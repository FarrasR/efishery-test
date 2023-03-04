package database

import (
	"github.com/go-rel/rel"
	"github.com/go-rel/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() rel.Adapter {
	adapter, err := sqlite3.Open("auth.db")
	if err != nil {
		panic(err)
	}

	return adapter
}
