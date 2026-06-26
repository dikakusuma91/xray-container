package repository

import (
	"database/sql"
	"time"

	"xray-post-test/models"
)

type RepositoriPemindaianXray struct {
	DB *sql.DB
}

func BuatRepositoriPemindaianXray(db *sql.DB) *RepositoriPemindaianXray {
	return &RepositoriPemindaianXray{DB: db}
}

func (r *RepositoriPemindaianXray) AmbilSemua() ([]models.PemindaianXray, error) {
	rows, err := r.DB.Query(
		"SELECT id, id_pengguna, nama_image, tanggal_pemindaian, deskripsi, path_gambar, prioritas, laporan, dibuat_pada, diperbarui_pada FROM pemindaian_xray ORDER BY dibuat_pada DESC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []models.PemindaianXray
	for rows.Next() {
		var h models.PemindaianXray
		if err := rows.Scan(&h.ID, &h.IDPengguna, &h.NamaImage, &h.TanggalPemindaian, &h.Deskripsi, &h.PathGambar, &h.Prioritas, &h.Laporan, &h.DibuatPada, &h.DiperbaruiPada); err != nil {
			return nil, err
		}
		hasil = append(hasil, h)
	}
	return hasil, nil
}

func (r *RepositoriPemindaianXray) AmbilBerdasarkanID(id string) (*models.PemindaianXray, error) {
	var h models.PemindaianXray
	row := r.DB.QueryRow(
		"SELECT id, id_pengguna, nama_image, tanggal_pemindaian, deskripsi, path_gambar, prioritas, laporan, dibuat_pada, diperbarui_pada FROM pemindaian_xray WHERE id = $1",
		id,
	)
	err := row.Scan(&h.ID, &h.IDPengguna, &h.NamaImage, &h.TanggalPemindaian, &h.Deskripsi, &h.PathGambar, &h.Prioritas, &h.Laporan, &h.DibuatPada, &h.DiperbaruiPada)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (r *RepositoriPemindaianXray) HitungSemuaPemindaian() (int, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM pemindaian_xray").Scan(&count)
	return count, err
}

func (r *RepositoriPemindaianXray) HitungBerdasarkanPrioritas(prioritas string) (int, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM pemindaian_xray WHERE prioritas = $1", prioritas).Scan(&count)
	return count, err
}

func (r *RepositoriPemindaianXray) AmbilPemindaianTerbaru(limit int) ([]models.PemindaianXray, error) {
	rows, err := r.DB.Query(
		"SELECT id, id_pengguna, nama_image, tanggal_pemindaian, deskripsi, path_gambar, prioritas, laporan, dibuat_pada, diperbarui_pada FROM pemindaian_xray ORDER BY dibuat_pada DESC LIMIT $1",
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []models.PemindaianXray
	for rows.Next() {
		var h models.PemindaianXray
		if err := rows.Scan(&h.ID, &h.IDPengguna, &h.NamaImage, &h.TanggalPemindaian, &h.Deskripsi, &h.PathGambar, &h.Prioritas, &h.Laporan, &h.DibuatPada, &h.DiperbaruiPada); err != nil {
			return nil, err
		}
		hasil = append(hasil, h)
	}
	return hasil, nil
}

func (r *RepositoriPemindaianXray) Simpan(h *models.PemindaianXray) error {
	_, err := r.DB.Exec(
		"INSERT INTO pemindaian_xray (id, id_pengguna, nama_image, tanggal_pemindaian, deskripsi, path_gambar, prioritas, laporan) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		h.ID, h.IDPengguna, h.NamaImage, h.TanggalPemindaian, h.Deskripsi, h.PathGambar, h.Prioritas, h.Laporan,
	)
	return err
}

func (r *RepositoriPemindaianXray) Perbarui(h *models.PemindaianXray) error {
	_, err := r.DB.Exec(
		"UPDATE pemindaian_xray SET nama_image=$1, tanggal_pemindaian=$2, deskripsi=$3, path_gambar=$4, laporan=$5, diperbarui_pada=$6 WHERE id=$7",
		h.NamaImage, h.TanggalPemindaian, h.Deskripsi, h.PathGambar, h.Laporan, time.Now(), h.ID,
	)
	return err
}

func (r *RepositoriPemindaianXray) Hapus(id string) error {
	_, err := r.DB.Exec("DELETE FROM pemindaian_xray WHERE id = $1", id)
	return err
}
