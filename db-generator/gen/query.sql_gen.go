// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package models

import (
	"context"
	"database/sql"
)

const createCategories = `-- name: CreateCategories :one
insert into categories(category_id, category_name, description)
values($1, $2, $3)
returning category_id
`

type CreateCategoriesParams struct {
	CategoryID   int16          `db:"category_id" json:"categoryId"`
	CategoryName sql.NullString `db:"category_name" json:"categoryName"`
	Description  sql.NullString `db:"description" json:"description"`
}

func (q *Queries) CreateCategories(ctx context.Context, arg CreateCategoriesParams) (int16, error) {
	row := q.db.QueryRowContext(ctx, createCategories, arg.CategoryID, arg.CategoryName, arg.Description)
	var category_id int16
	err := row.Scan(&category_id)
	return category_id, err
}

const createTransaction = `-- name: CreateTransaction :one
insert into transaction(transact_id, type_transact, description, nominal, time_of_transaction, category_id)
values($1, $2, $3, $4, $5, $6)
returning transact_id
`

type CreateTransactionParams struct {
	TransactID        string         `db:"transact_id" json:"transactId"`
	TypeTransact      sql.NullString `db:"type_transact" json:"typeTransact"`
	Description       sql.NullString `db:"description" json:"description"`
	Nominal           sql.NullString `db:"nominal" json:"nominal"`
	TimeOfTransaction sql.NullTime   `db:"time_of_transaction" json:"timeOfTransaction"`
	CategoryID        sql.NullInt16  `db:"category_id" json:"categoryId"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.TransactID,
		arg.TypeTransact,
		arg.Description,
		arg.Nominal,
		arg.TimeOfTransaction,
		arg.CategoryID,
	)
	var transact_id string
	err := row.Scan(&transact_id)
	return transact_id, err
}

const deleteCategories = `-- name: DeleteCategories :exec
delete from categories
where category_id = $1
`

func (q *Queries) DeleteCategories(ctx context.Context, categoryID int16) error {
	_, err := q.db.ExecContext(ctx, deleteCategories, categoryID)
	return err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
delete from transaction
where transact_id = $1
`

func (q *Queries) DeleteTransaction(ctx context.Context, transactID string) error {
	_, err := q.db.ExecContext(ctx, deleteTransaction, transactID)
	return err
}

const getCategories = `-- name: GetCategories :one
select category_id, category_name, description from categories
where category_id = $1
`

func (q *Queries) GetCategories(ctx context.Context, categoryID int16) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategories, categoryID)
	var i Category
	err := row.Scan(&i.CategoryID, &i.CategoryName, &i.Description)
	return i, err
}

const getTransaction = `-- name: GetTransaction :one
select transact_id, type_transact, description, nominal, time_of_transaction, category_id from transaction
where transact_id = $1
`

func (q *Queries) GetTransaction(ctx context.Context, transactID string) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransaction, transactID)
	var i Transaction
	err := row.Scan(
		&i.TransactID,
		&i.TypeTransact,
		&i.Description,
		&i.Nominal,
		&i.TimeOfTransaction,
		&i.CategoryID,
	)
	return i, err
}

const listCategories = `-- name: ListCategories :many
Select category_id, category_name, description from categories
order by category_id
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.CategoryID, &i.CategoryName, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransaction = `-- name: ListTransaction :many
select transact_id, type_transact, description, nominal, time_of_transaction, category_id from transaction
order by category_id
`

func (q *Queries) ListTransaction(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransaction)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TransactID,
			&i.TypeTransact,
			&i.Description,
			&i.Nominal,
			&i.TimeOfTransaction,
			&i.CategoryID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategories = `-- name: UpdateCategories :exec
update categories
    set category_id = $1,
    category_name = $2,
    description = $3
where category_id = $1
`

type UpdateCategoriesParams struct {
	CategoryID   int16          `db:"category_id" json:"categoryId"`
	CategoryName sql.NullString `db:"category_name" json:"categoryName"`
	Description  sql.NullString `db:"description" json:"description"`
}

func (q *Queries) UpdateCategories(ctx context.Context, arg UpdateCategoriesParams) error {
	_, err := q.db.ExecContext(ctx, updateCategories, arg.CategoryID, arg.CategoryName, arg.Description)
	return err
}

const updateTransaction = `-- name: UpdateTransaction :exec
update transaction
    set transact_id = $1,
    type_transact = $2,
    description = $3,
    nominal = $4,
    time_of_transaction = $5,
    category_id = $6
where transact_id = $1
`

type UpdateTransactionParams struct {
	TransactID        string         `db:"transact_id" json:"transactId"`
	TypeTransact      sql.NullString `db:"type_transact" json:"typeTransact"`
	Description       sql.NullString `db:"description" json:"description"`
	Nominal           sql.NullString `db:"nominal" json:"nominal"`
	TimeOfTransaction sql.NullTime   `db:"time_of_transaction" json:"timeOfTransaction"`
	CategoryID        sql.NullInt16  `db:"category_id" json:"categoryId"`
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) error {
	_, err := q.db.ExecContext(ctx, updateTransaction,
		arg.TransactID,
		arg.TypeTransact,
		arg.Description,
		arg.Nominal,
		arg.TimeOfTransaction,
		arg.CategoryID,
	)
	return err
}