-- name: CreateBook :one
INSERT INTO books (id, created_at, updated_at, title, volume, category, author, published_at, publisher, finished, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetBooks :many
SELECT * FROM books;

-- name: UpdateBook :exec
UPDATE books
SET title = $3, volume = $4, category = $5, author = $6, published_at = $7, publisher = $8, finished = $9, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1 AND user_id = $2;