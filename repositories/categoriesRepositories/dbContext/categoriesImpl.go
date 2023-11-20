package dbContext

import (
	"context"
	models "technopartnertest_go/models"
)

const getCategories = `-- name: GetCategories :one
select category_id, category_name, description from categories
where category_id = $1
`

func (q *Queries) GetCategories(ctx context.Context, categoryID int16) (models.Category, error) {
	row := q.db.QueryRowContext(ctx, getCategories, categoryID)
	var i models.Category
	err := row.Scan(&i.CategoryID, &i.CategoryName, &i.Description)
	return i, err
}
