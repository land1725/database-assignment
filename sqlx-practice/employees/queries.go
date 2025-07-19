package main

import "github.com/jmoiron/sqlx"

func getEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM gorm.employees where department='技术部'")
	return employees, err
}

func getHighestEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	err := db.Get(&employee, "select * from gorm.employees order by salary desc limit 1")
	return employee, err
}
