-- name: GetQuangByEmail :one
SELECT * FROM quang_test_v1_goose WHERE name = $1 LIMIT 1;

-- name: UpdateQuang :exec
UPDATE quang_test_v1_goose 
    set name = $2, 
    age = $3
WHERE id = $1;

-- name: CreateQuang :one
INSERT INTO quang_test_v1_goose (
  name, age, sex
) VALUES (
  $1, $2, $3
)
RETURNING *;