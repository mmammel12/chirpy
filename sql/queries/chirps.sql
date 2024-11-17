-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $2,
    $1
)
RETURNING *;

-- name: DeleteChirps :exec
DELETE FROM chirps;

-- name: ListChirps :many
SELECT * FROM chirps
ORDER BY created_at;

-- name: GetChirpById :one
SELECT * FROM chirps
WHERE id = $1;
