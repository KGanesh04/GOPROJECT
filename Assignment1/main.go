package main

import (
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e *Employee) UpdateSalary(salary float64) {
	e.Salary = salary
}
func AddEmployee(emp *Employee, db map[int]*Employee) {
	db[emp.ID] = emp
}

func RemoveEmployee(id int, db map[int]*Employee) {
	delete(db, id)
}

func main() {

	E1 := Employee{ID: 1, Name: "Alice", Salary: 50000}
	E2 := Employee{ID: 2, Name: "Bob", Salary: 60000}
	E3 := Employee{ID: 3, Name: "Charlie", Salary: 70000}

	fmt.Println("Employee 1 salary:", E1.Salary)
	fmt.Println("Employee 2 salary:", E2.Salary)
	fmt.Println("Employee 3 salary:", E3.Salary)

	var salary float64 = 70000

	E1.UpdateSalary(salary)
	fmt.Println("Employee 1 after salary update:", E1.Salary)

	employeeDB := map[int]*Employee{
		1: {ID: 1, Name: "Alice", Salary: 50000},
		2: {ID: 2, Name: "Bob", Salary: 60000},
	}

	AddEmployee(&Employee{ID: 3, Name: "Charlie", Salary: 70000}, employeeDB)
	fmt.Println("Employee 3 added:", employeeDB[3])

	RemoveEmployee(2, employeeDB)
	fmt.Println("Employee 2 removed:", employeeDB[2])

	for id, emp := range employeeDB {
		fmt.Println("Employee ID:", id, "Name:", emp.Name, "Salary:", emp.Salary)
	}
}
