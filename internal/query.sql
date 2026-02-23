-- name: CreateEntry :one
INSERT INTO entries (
	source,
	target,
	amount,
	description
) VALUES (?, ?, ?, ?) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries WHERE id = (?);

-- name: GetAllEntries :many
SELECT * FROM entries;

-- name: UpdateEntry :one
UPDATE entries
SET source = ?,
target = ?,
amount = ?,
description = ?
WHERE id = (?)
RETURNING *;


-- name: DeleteEntry :one
DELETE FROM entries
WHERE id = (?)
RETURNING *;
