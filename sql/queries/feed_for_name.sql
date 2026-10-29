-- name: GetFeedForName :one
Select * FROM feeds
WHERE name = $1 LIMIT 1;
