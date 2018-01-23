package sql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDBConnection(t *testing.T) {
	db, err := sql.Open("mysql", "test:test@/test_dat")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r, err := db.Query("show tables;")
	//	r, err := db.Query("select * from nothing")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	fmt.Printf("row: %+v\n", r)

	for r.Next() {
		fmt.Printf("row: %+v\n", r)
	}

}
