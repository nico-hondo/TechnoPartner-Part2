package models

import (
	"database/sql"
)

type Category struct {
	CategoryID   int16          `db:"category_id" json:"categoryId"`
	CategoryName sql.NullString `db:"category_name" json:"categoryName"`
	Description  sql.NullString `db:"description" json:"description"`
}

type Transaction struct {
	TransactID        string         `db:"transact_id" json:"transactId"`
	TypeTransact      sql.NullString `db:"type_transact" json:"typeTransact"`
	Description       sql.NullString `db:"description" json:"description"`
	Nominal           sql.NullString `db:"nominal" json:"nominal"`
	TimeOfTransaction sql.NullTime   `db:"time_of_transaction" json:"timeOfTransaction"`
	CategoryID        sql.NullInt16  `db:"category_id" json:"categoryId"`
}
