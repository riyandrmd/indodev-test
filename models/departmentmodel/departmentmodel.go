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

func Insert(department entities.Departments) bool {
	result, err := config.DB.Exec(`
		INSERT INTO Departments (DepartmentCode, DepartmentName)
		VALUES (?, ?)`,
		department.DepartmentCode,
		department.DepartmentName,
	)

	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}

func Detail(id string) entities.Departments {
	row := config.DB.QueryRow(`SELECT * FROM Departments WHERE DepartmentCode = ?`, id)

	var department entities.Departments
	if err := row.Scan(&department.DepartmentCode, &department.DepartmentName); err != nil {
		panic(err.Error())
	}

	return department
}

func Update(id string, department entities.Departments) bool {
	query, err := config.DB.Exec("UPDATE Departments SET DepartmentName = ? WHERE DepartmentCode = ?", department.DepartmentName, id)
	if err != err {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id string) error {
	_, err := config.DB.Exec("DELETE FROM Departments WHERE DepartmentCode = ?", id)
	return err
}
