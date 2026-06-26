package models

import "time"

type PemindaianXray struct {
	ID                string    `json:"id"`
	IDPengguna        string    `json:"id_pengguna"`
	NamaImage         string    `json:"nama_image"`
	TanggalPemindaian time.Time `json:"tanggal_pemindaian"`
	Deskripsi         string    `json:"deskripsi"`
	PathGambar        string    `json:"path_gambar"`
	Prioritas         string    `json:"prioritas"`
	Laporan           string    `json:"laporan"`
	DibuatPada        time.Time `json:"dibuat_pada"`
	DiperbaruiPada    time.Time `json:"diperbarui_pada"`
}
