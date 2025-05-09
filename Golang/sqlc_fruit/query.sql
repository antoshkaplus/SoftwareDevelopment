-- name: ListFruit :many
SELECT * FROM fruit;

-- name: CreateFruit :one
INSERT INTO fruit (
  name, created
) VALUES (
  ?, ?
)
RETURNING *;

-- name: DeleteFruit :exec
DELETE FROM fruit
WHERE name = ?;