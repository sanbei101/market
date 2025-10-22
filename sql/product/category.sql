-- name: CreateCategory :one
INSERT INTO category (
    name, 
    parent_id, 
    spec_template, 
    sort
) VALUES (
    @name, @parent_id, @spec_template, @sort
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM category WHERE id = @id;

-- name: ListCategories :many
SELECT * FROM category 
ORDER BY sort ASC, created_at DESC;

-- name: ListCategoriesByParent :many
SELECT * FROM category 
WHERE parent_id = @parent_id
ORDER BY sort ASC, created_at DESC;

-- name: UpdateCategory :one
UPDATE category 
SET 
    name = @name,
    parent_id = @parent_id,
    spec_template = @spec_template,
    sort = @sort
WHERE id = @id
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category WHERE id = @id;