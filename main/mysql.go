package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"reflect"
)


func main()  {
	db, err := sql.Open("mysql", "root:1234567qq@@tcp(127.0.0.1:3306)/?charset=utf8")
	if err !=nil{
		fmt.Println(err)
	}
	db.Query("drop database if exists tmpdb")
	db.Query("create database tmpdb")
	db.Query("create table tmpdb.tmptab(c1 int, c2 varchar(20), c3 varchar(20))")
	db.Query("insert into tmpdb.tmptab values(101, '姓名1', 'address1'), (102, '姓名2', 'address2'), (103, '姓名3', 'address3'), (104, '姓名4', 'address4')")
	query, err := db.Query("select * from tmpdb.tmptab")
	if err !=nil{
		fmt.Println(err)
	}
	v := reflect.ValueOf(query)
	fmt.Println(v)

}
