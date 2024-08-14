-- name: AddQuest :one
INSERT INTO quests (id, owner, email, name, description, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, owner, email, name, description, status;

-- name: GetAssignedQuests :many 
SELECT id, owner, email, name, description, status
FROM quests
WHERE email = $1;

-- name: CompleteQuest :exec
UPDATE quests
SET status = 'completed'
WHERE id = $1;

-- name: GetQuest :one
SELECT id, owner, email, name, description, status
FROM quests
WHERE id = $1;