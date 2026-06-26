package services

import (
	"xray-post-test/models"
	"xray-post-test/repository"
)

type LayananDasbor struct {
	RepoPengguna    *repository.RepositoriPengguna
	RepoPemindaian  *repository.RepositoriPemindaianXray
}

func BuatLayananDasbor(repoPengguna *repository.RepositoriPengguna, repoPemindaian *repository.RepositoriPemindaianXray) *LayananDasbor {
	return &LayananDasbor{RepoPengguna: repoPengguna, RepoPemindaian: repoPemindaian}
}

type RingkasanDasbor struct {
	TotalPengguna      int                       `json:"total_pengguna"`
	TotalPemindaian    int                       `json:"total_pemindaian"`
	PrioritasTinggi    int                       `json:"prioritas_tinggi"`
	PrioritasSedang    int                       `json:"prioritas_sedang"`
	PrioritasRendah    int                       `json:"prioritas_rendah"`
	PemindaianTerbaru  []models.PemindaianXray   `json:"pemindaian_terbaru"`
	PenggunaTerbaru    []models.Pengguna         `json:"pengguna_terbaru"`
}

func (s *LayananDasbor) AmbilRingkasan() (*RingkasanDasbor, error) {
	totalPengguna, err := s.RepoPengguna.HitungSemuaPengguna()
	if err != nil {
		return nil, err
	}

	totalPemindaian, err := s.RepoPemindaian.HitungSemuaPemindaian()
	if err != nil {
		return nil, err
	}

	prioritasTinggi, _ := s.RepoPemindaian.HitungBerdasarkanPrioritas("tinggi")
	prioritasSedang, _ := s.RepoPemindaian.HitungBerdasarkanPrioritas("sedang")
	prioritasRendah, _ := s.RepoPemindaian.HitungBerdasarkanPrioritas("rendah")

	pemindaianTerbaru, _ := s.RepoPemindaian.AmbilPemindaianTerbaru(5)
	penggunaTerbaru, _ := s.RepoPengguna.AmbilPenggunaTerbaru(5)

	return &RingkasanDasbor{
		TotalPengguna:     totalPengguna,
		TotalPemindaian:   totalPemindaian,
		PrioritasTinggi:   prioritasTinggi,
		PrioritasSedang:   prioritasSedang,
		PrioritasRendah:   prioritasRendah,
		PemindaianTerbaru: pemindaianTerbaru,
		PenggunaTerbaru:   penggunaTerbaru,
	}, nil
}
