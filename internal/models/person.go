package models

import "database/sql"

type Person struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Surname     string         `json:"surname"`
	Patronymic  sql.NullString `json:"patronymic,omitempty"`
	Age         sql.NullInt32  `json:"age,omitempty"`
	Gender      sql.NullString `json:"gender,omitempty"`
	Nationality sql.NullString `json:"nationality,omitempty"`
}
