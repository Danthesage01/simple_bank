-- name: CreateTransfer :one
INSERT INTO tranfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTranfer :one
SELECT * FROM tranfers
WHERE id = $1 LIMIT 1;

-- name: ListTranfers :many
SELECT * FROM tranfers
WHERE 
  from_account_id = $1 OR
  to_account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

