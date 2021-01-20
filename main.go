package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type employees struct {
	Status string `json:"status"`
	Data   []struct {
		ID             string `json:"id"`
		EmployeeName   string `json:"employee_name"`
		EmployeeSalary string `json:"employee_salary"`
		EmployeeAge    string `json:"employee_age"`
		ProfileImage   string `json:"profile_image"`
	} `json:"data"`
}

type SalaryCalculator interface {
	CalculateSalary() int
}

func (e employees) CalculateSalary() int {
	sum := 0
	for i := range e.Data {
		x := e.Data[i].EmployeeSalary
		salary, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println(err)
		}
		sum += salary
	}
	return sum
}
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Println("Total Salary ", expense)
}

func main() {
	resp, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	var employee employees
	err = json.Unmarshal([]byte(resp), &employee)
	if err != nil {
		fmt.Println(err)
	}

	salary := []SalaryCalculator{employee}
	totalExpense(salary)
}
