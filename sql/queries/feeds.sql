-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds
ORDER BY name;

-- name: GetFeed :one
Select * FROM feeds
WHERE id = $1 LIMIT 1;

-- name: GetFeedForName :one
Select * FROM feeds
WHERE name = $1 LIMIT 1;

-- name: GetFeedForURL :one
Select * FROM feeds
WHERE url = $1 LIMIT 1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $2,
    updated_at = $2
WHERE id = $1;
