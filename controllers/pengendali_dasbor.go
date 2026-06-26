package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xray-post-test/services"
)

type PengendaliDasbor struct {
	Layanan *services.LayananDasbor
}

func BuatPengendaliDasbor(layanan *services.LayananDasbor) *PengendaliDasbor {
	return &PengendaliDasbor{Layanan: layanan}
}

// Ringkasan godoc
//
//	@Summary		Ringkasan dashboard
//	@Tags			Dasbor
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	services.RingkasanDasbor
//	@Failure		500	{object}	map[string]string
//	@Router			/api/v1/dasbor/ringkasan [get]
func (p *PengendaliDasbor) Ringkasan(c *gin.Context) {
	data, err := p.Layanan.AmbilRingkasan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
