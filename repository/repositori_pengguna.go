package repository

import (
	"database/sql"

	"xray-post-test/models"
)

type RepositoriPengguna struct {
	DB *sql.DB
}

func BuatRepositoriPengguna(db *sql.DB) *RepositoriPengguna {
	return &RepositoriPengguna{DB: db}
}

func (r *RepositoriPengguna) CariBerdasarkanEmail(email string) (*models.Pengguna, error) {
	var p models.Pengguna
	row := r.DB.QueryRow(
		"SELECT id, email, kata_sandi, nama_lengkap, dibuat_pada FROM pengguna WHERE email = $1",
		email,
	)
	err := row.Scan(&p.ID, &p.Email, &p.KataSandi, &p.NamaLengkap, &p.DibuatPada)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *RepositoriPengguna) Simpan(p *models.Pengguna) error {
	_, err := r.DB.Exec(
		"INSERT INTO pengguna (id, email, kata_sandi, nama_lengkap) VALUES ($1, $2, $3, $4)",
		p.ID, p.Email, p.KataSandi, p.NamaLengkap,
	)
	return err
}
