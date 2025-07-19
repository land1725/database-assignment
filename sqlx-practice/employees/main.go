package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
	"github.com/jmoiron/sqlx"
)

func main() {
	// 连接字符串格式: "username:password@tcp(host:port)/dbname?parseTime=true"
	db, err := sqlx.Connect("mysql", "root:land1725@tcp(127.0.0.1:3306)/gorm?parseTime=true")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer db.Close()

	fmt.Println("成功连接到数据库")
	employees, err := getEmployees(db) // 或者使用 employee, err := getEmployees(db)
	if err != nil {
		fmt.Println("获取员工数据失败:", err)
		return
	} else {
		fmt.Println(employees)
	}
	employee, err := getHighestEmployee(db)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(employee)
	}
}
