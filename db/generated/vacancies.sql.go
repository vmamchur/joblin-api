// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: vacancies.sql

package generated

import (
	"context"
	"database/sql"
)

const createVacancy = `-- name: CreateVacancy :one
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
RETURNING id, created_at, updated_at, title, company_name, url
`

type CreateVacancyParams struct {
	Title       string         `json:"title"`
	CompanyName sql.NullString `json:"company_name"`
	Url         string         `json:"url"`
}

func (q *Queries) CreateVacancy(ctx context.Context, arg CreateVacancyParams) (Vacancy, error) {
	row := q.db.QueryRowContext(ctx, createVacancy, arg.Title, arg.CompanyName, arg.Url)
	var i Vacancy
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.CompanyName,
		&i.Url,
	)
	return i, err
}

const getAllVacancies = `-- name: GetAllVacancies :many
SELECT id, created_at, updated_at, title, company_name, url FROM vacancies
`

func (q *Queries) GetAllVacancies(ctx context.Context) ([]Vacancy, error) {
	rows, err := q.db.QueryContext(ctx, getAllVacancies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vacancy{}
	for rows.Next() {
		var i Vacancy
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.CompanyName,
			&i.Url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
