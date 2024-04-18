package divisionsmodel

import (
	"indodev-test/config"
	"indodev-test/entities"
)

func GetAll() []entities.Divisions {
	rows, err := config.DB.Query("SELECT * FROM divisions")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var divisions []entities.Divisions

	for rows.Next() {
		var division entities.Divisions
		if err := rows.Scan(&division.DivisionId, &division.DivisionName, &division.DepartmentId); err != nil {
			panic(err)
		}

		divisions = append(divisions, division)
	}

	return divisions
}

func Insert(division entities.Divisions) bool {
	result, err := config.DB.Exec(`
		INSERT INTO Divisions (DivisionId, DivisionName, DepartmentId)
		VALUES (?, ?, ?)`,
		division.DivisionId,
		division.DivisionName,
		division.DepartmentId,
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

func Detail(id string) entities.Divisions {
	row := config.DB.QueryRow(`SELECT * FROM Divisions WHERE DivisionId = ?`, id)

	var division entities.Divisions
	if err := row.Scan(&division.DivisionId, &division.DivisionName, &division.DepartmentId); err != nil {
		panic(err.Error())
	}

	return division
}

func Update(id string, division entities.Divisions) bool {
	query, err := config.DB.Exec("UPDATE Divisions SET DivisionName = ? WHERE DivisionId = ?", division.DivisionName, id)
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
	_, err := config.DB.Exec("DELETE FROM Divisions WHERE DivisionId = ?", id)
	return err
}
