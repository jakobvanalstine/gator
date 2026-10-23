-- name: GetUserForID :one
Select * FROM users
WHERE id = $1 LIMIT 1;
