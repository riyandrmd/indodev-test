package departmentmodel

import (
	"indodev-test/config"
	"indodev-test/entities"
)

func GetAll() []entities.Departments {
	rows, err := config.DB.Query("SELECT * FROM departments")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var departments []entities.Departments

	for rows.Next() {
		var department entities.Departments
		if err := rows.Scan(&department.DepartmentCode, &department.DepartmentName); err != nil {
			panic(err)
		}

		departments = append(departments, department)
	}

	return departments
}
