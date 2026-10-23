-- name: GetFeedForURL :one
Select * FROM feeds
WHERE url = $1 LIMIT 1;
