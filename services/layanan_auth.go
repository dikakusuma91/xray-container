package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"xray-post-test/config"
	"xray-post-test/models"
	"xray-post-test/repository"
)

type LayananAuth struct {
	Repo *repository.RepositoriPengguna
	Cfg  *config.Konfigurasi
}

func BuatLayananAuth(repo *repository.RepositoriPengguna, cfg *config.Konfigurasi) *LayananAuth {
	return &LayananAuth{Repo: repo, Cfg: cfg}
}

type KlaimJWT struct {
	IDPengguna  string `json:"id_pengguna"`
	Email       string `json:"email"`
	NamaLengkap string `json:"nama_lengkap"`
	jwt.RegisteredClaims
}

func (s *LayananAuth) Daftar(email, kataSandi, namaLengkap string) (*models.Pengguna, string, error) {
	ada, _ := s.Repo.CariBerdasarkanEmail(email)
	if ada != nil {
		return nil, "", errors.New("email sudah terdaftar")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(kataSandi), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	pengguna := &models.Pengguna{
		ID:          uuid.New().String(),
		Email:       email,
		KataSandi:   string(hash),
		NamaLengkap: namaLengkap,
		DibuatPada:  time.Now(),
	}

	if err := s.Repo.Simpan(pengguna); err != nil {
		return nil, "", err
	}

	token, err := s.buatToken(pengguna)
	if err != nil {
		return nil, "", err
	}

	return pengguna, token, nil
}

func (s *LayananAuth) Masuk(email, kataSandi string) (*models.Pengguna, string, error) {
	pengguna, err := s.Repo.CariBerdasarkanEmail(email)
	if err != nil {
		return nil, "", errors.New("email atau kata sandi salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(pengguna.KataSandi), []byte(kataSandi)); err != nil {
		return nil, "", errors.New("email atau kata sandi salah")
	}

	token, err := s.buatToken(pengguna)
	if err != nil {
		return nil, "", err
	}

	return pengguna, token, nil
}

func (s *LayananAuth) buatToken(p *models.Pengguna) (string, error) {
	klaim := KlaimJWT{
		p.ID,
		p.Email,
		p.NamaLengkap,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, klaim)
	return token.SignedString([]byte(s.Cfg.JwtRahasia))
}
