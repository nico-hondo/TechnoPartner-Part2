-- name: GetCategories :one
select * from categories
where category_id = $1;

-- name: ListCategories :many
Select * from categories
order by category_id;

-- name: CreateCategories :one
insert into categories(category_id, category_name, description)
values($1, $2, $3)
returning category_id;

-- name: UpdateCategories :exec
update categories
    set category_id = $1,
    category_name = $2,
    description = $3
where category_id = $1;

-- name: DeleteCategories :exec
delete from categories
where category_id = $1;


-- name: GetTransaction :one
select * from transaction
where transact_id = $1;

-- name: ListTransaction :many
select * from transaction
order by category_id;

-- name: CreateTransaction :one
insert into transaction(transact_id, type_transact, description, nominal, time_of_transaction, category_id)
values($1, $2, $3, $4, $5, $6)
returning transact_id;

-- name: UpdateTransaction :exec
update transaction
    set transact_id = $1,
    type_transact = $2,
    description = $3,
    nominal = $4,
    time_of_transaction = $5,
    category_id = $6
where transact_id = $1;

-- name: DeleteTransaction :exec
delete from transaction
where transact_id = $1;