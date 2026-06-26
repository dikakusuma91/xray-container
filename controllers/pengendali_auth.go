package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xray-post-test/services"
)

type PengendaliAuth struct {
	Layanan *services.LayananAuth
}

func BuatPengendaliAuth(layanan *services.LayananAuth) *PengendaliAuth {
	return &PengendaliAuth{Layanan: layanan}
}

type PermintaanDaftar struct {
	Email       string `json:"email" binding:"required,email"`
	KataSandi   string `json:"kata_sandi" binding:"required,min=6"`
	NamaLengkap string `json:"nama_lengkap" binding:"required"`
}

type PermintaanMasuk struct {
	Email     string `json:"email" binding:"required,email"`
	KataSandi string `json:"kata_sandi" binding:"required"`
}

// Daftar godoc
//
//	@Summary		Daftar pengguna baru
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		PermintaanDaftar	true	"Data pendaftaran"
//	@Success		201		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]string
//	@Router			/api/v1/auth/daftar [post]
func (p *PengendaliAuth) Daftar(c *gin.Context) {
	var req PermintaanDaftar
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pengguna, token, err := p.Layanan.Daftar(req.Email, req.KataSandi, req.NamaLengkap)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"pengguna": pengguna,
		"token":    token,
	})
}

// Masuk godoc
//
//	@Summary		Masuk dan dapatkan token JWT
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		PermintaanMasuk	true	"Data login"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		401		{object}	map[string]string
//	@Router			/api/v1/auth/masuk [post]
func (p *PengendaliAuth) Masuk(c *gin.Context) {
	var req PermintaanMasuk
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pengguna, token, err := p.Layanan.Masuk(req.Email, req.KataSandi)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pengguna": pengguna,
		"token":    token,
	})
}
