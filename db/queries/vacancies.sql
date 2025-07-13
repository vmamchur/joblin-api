-- name: CreateVacancy :one
INSERT INTO vacancies (id, created_at, updated_at, title, company_name, url)
VALUES (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2,
	$3
)
ON CONFLICT (url) DO NOTHING
RETURNING *;

-- name: GetVacancies :many
SELECT *
FROM vacancies
WHERE (
	$1::text IS NULL
	OR company_name ILIKE '%' || $1::text || '%'
	OR title ILIKE '%' || $1::text || '%'
)
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

