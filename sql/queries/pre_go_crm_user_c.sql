-- name: GetUserByEmail :one
SELECT * FROM user_v1_goose WHERE name = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE user_v1_goose 
    set name = $2, 
    age = $3
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO user_v1_goose (
  name, age, sex
) VALUES (
  $1, $2, $3
)
RETURNING *;