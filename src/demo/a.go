package demo

import (
	"database/sql"
)

type A struct {
	Db  *sql.DB `di:"db"`
	Db1 *sql.DB `di:"db"`
	B   *B      `di:"b,prototype"`
	B1  *B      `di:"b,prototype"`
}

func NewA() *A {
	return &A{}
}

func (p *A) Version() (string, error) {
	rows, err := p.Db.Query("SELECT VERSION() as version")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var version string
	if rows.Next() {
		if err := rows.Scan(&version); err != nil {
			return "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	return version, nil
}
