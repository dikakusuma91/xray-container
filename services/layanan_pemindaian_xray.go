package services

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"xray-post-test/models"
	"xray-post-test/repository"
)

type LayananPemindaianXray struct {
	Repo *repository.RepositoriPemindaianXray
}

func BuatLayananPemindaianXray(repo *repository.RepositoriPemindaianXray) *LayananPemindaianXray {
	return &LayananPemindaianXray{Repo: repo}
}

func (s *LayananPemindaianXray) AmbilSemua() ([]models.PemindaianXray, error) {
	return s.Repo.AmbilSemua()
}

func (s *LayananPemindaianXray) AmbilBerdasarkanID(id string) (*models.PemindaianXray, error) {
	return s.Repo.AmbilBerdasarkanID(id)
}

func (s *LayananPemindaianXray) BuatBaru(idPengguna, namaImage, deskripsi, pathGambar, laporan string) (*models.PemindaianXray, error) {
	h := &models.PemindaianXray{
		ID:                uuid.New().String(),
		IDPengguna:        idPengguna,
		NamaImage:         namaImage,
		TanggalPemindaian: time.Now(),
		Deskripsi:         deskripsi,
		PathGambar:        pathGambar,
		Prioritas:         prioritasAcak(),
		Laporan:           laporan,
		DibuatPada:        time.Now(),
		DiperbaruiPada:    time.Now(),
	}

	if err := s.Repo.Simpan(h); err != nil {
		return nil, err
	}
	return h, nil
}

func (s *LayananPemindaianXray) Perbarui(id, namaImage, deskripsi, pathGambar, laporan string, tanggalPemindaian time.Time) (*models.PemindaianXray, error) {
	ada, err := s.Repo.AmbilBerdasarkanID(id)
	if err != nil {
		return nil, err
	}

	ada.NamaImage = namaImage
	ada.Deskripsi = deskripsi
	ada.PathGambar = pathGambar
	ada.Laporan = laporan
	ada.TanggalPemindaian = tanggalPemindaian

	if err := s.Repo.Perbarui(ada); err != nil {
		return nil, err
	}
	return ada, nil
}

func (s *LayananPemindaianXray) Hapus(id string) error {
	return s.Repo.Hapus(id)
}

func prioritasAcak() string {
	n := rand.Intn(100)
	switch {
	case n < 50:
		return "rendah"
	case n < 80:
		return "sedang"
	default:
		return "tinggi"
	}
}
