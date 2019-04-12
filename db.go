package main

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/goracle.v2"
)

func DB() {
	db, err := sql.Open("goracle", "user/pass@127.0.0.1:1521/oracledb")
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from durl")
	if err != nil {
		fmt.Println("error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	var thedata string
	for row.Next() {
		rows.Scan(&thedata)
	}
	fmt.Println("The date is: %s\n", thedate)

}
