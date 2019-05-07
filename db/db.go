package db

import (
	"book/conf"
	"database/sql"
	"fmt"
	"github.com/mysql"
)

var (
	db        *sql.DB
	isConnect bool
	abc mysql.MySQLDriver
)

func init() {
	var err error
	db, err = connect()
	if err != nil {
		e(err)
	}
	isConnect = true;
}

func connect() (db *sql.DB, error error) {
	db, err := sql.Open("mysql", conf.DB_Driver)
	return db, err
}

func e(err error) {
	panic(err)
	fmt.Println(err)
}

func Query(sql string) *sql.Rows {
	rows, err := db.Query(sql)
	if err != nil {
		e(err)
	}
	return rows
}

func exec(sql string, args ...interface{}) sql.Result {
	if !isConnect {
		_, _ = connect()
	}

	//st, err := db.Prepare(sql)
	//if err != nil {
	//	e(err)
	//}
	fmt.Println("sql",sql,args)

	res, err := db.Exec(sql)

	if err != nil {
		e(err)
	}
	return res
}

//update table set field1=? where id=?
func Insert(sql string, args ...interface{}) int64 {
	res := exec(sql, args...)
	id, err := res.LastInsertId()
	if err != nil {
		e(err)
	}
	return id
}

func Update(sql string, args ...interface{}) int64 {
	res := exec(sql, args...)
	rows, err := res.RowsAffected()
	if err != nil {
		e(err)
	}
	return rows
}

func Delete(sql string, args ...interface{}) int64 {
	res := exec(sql, args...)
	rows, err := res.RowsAffected()
	if err != nil {
		e(err)
	}
	return rows
}
