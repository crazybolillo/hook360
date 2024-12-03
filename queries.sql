-- name: InsertEvent :exec
INSERT INTO event (payload) VALUES ($1);
