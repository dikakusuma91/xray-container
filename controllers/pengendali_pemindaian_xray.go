package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"xray-post-test/config"
	"xray-post-test/services"
)

type PengendaliPemindaianXray struct {
	Layanan *services.LayananPemindaianXray
	Cfg     *config.Konfigurasi
}

func BuatPengendaliPemindaianXray(layanan *services.LayananPemindaianXray, cfg *config.Konfigurasi) *PengendaliPemindaianXray {
	return &PengendaliPemindaianXray{Layanan: layanan, Cfg: cfg}
}

// AmbilSemua godoc
//
//	@Summary		Ambil semua pemindaian xray
//	@Tags			PemindaianXray
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{array}		models.PemindaianXray
//	@Failure		500	{object}	map[string]string
//	@Router			/api/v1/pemindaian-xray [get]
func (p *PengendaliPemindaianXray) AmbilSemua(c *gin.Context) {
	data, err := p.Layanan.AmbilSemua()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// AmbilBerdasarkanID godoc
//
//	@Summary		Ambil pemindaian xray berdasarkan ID
//	@Tags			PemindaianXray
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"ID Pemindaian"
//	@Success		200	{object}	models.PemindaianXray
//	@Failure		404	{object}	map[string]string
//	@Router			/api/v1/pemindaian-xray/{id} [get]
func (p *PengendaliPemindaianXray) AmbilBerdasarkanID(c *gin.Context) {
	id := c.Param("id")
	data, err := p.Layanan.AmbilBerdasarkanID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, data)
}

// BuatBaru godoc
//
//	@Summary		Buat pemindaian xray baru
//	@Tags			PemindaianXray
//	@Accept			multipart/form-data
//	@Produce		json
//	@Security		BearerAuth
//	@Param			nama_image			formData	string	true	"Nama image container"
//	@Param			tanggal_pemindaian	formData	string	false	"Tanggal pemindaian (RFC3339)"
//	@Param			deskripsi			formData	string	false	"Deskripsi"
//	@Param			laporan				formData	string	false	"Laporan hasil pemindaian"
//	@Param			gambar				formData	file	false	"File gambar hasil scan"
//	@Success		201					{object}	models.PemindaianXray
//	@Failure		400					{object}	map[string]string
//	@Router			/api/v1/pemindaian-xray [post]
func (p *PengendaliPemindaianXray) BuatBaru(c *gin.Context) {
	namaImage := c.PostForm("nama_image")
	if namaImage == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama_image wajib diisi"})
		return
	}

	tglStr := c.DefaultPostForm("tanggal_pemindaian", time.Now().Format(time.RFC3339))
	tgl, err := time.Parse(time.RFC3339, tglStr)
	if err != nil {
		tgl = time.Now()
	}

	deskripsi := c.PostForm("deskripsi")
	laporan := c.PostForm("laporan")

	pathGambar := ""
	file, err := c.FormFile("gambar")
	if err == nil {
		ext := filepath.Ext(file.Filename)
		namaFile := uuid.New().String() + ext
		pathGambar = filepath.Join(p.Cfg.FolderGambar, namaFile)

		if err := c.SaveUploadedFile(file, pathGambar); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan gambar"})
			return
		}
	}

	idPengguna, _ := c.Get("id_pengguna")
	data, err := p.Layanan.BuatBaru(
		fmt.Sprintf("%v", idPengguna),
		namaImage, deskripsi, pathGambar, laporan,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data.TanggalPemindaian = tgl
	c.JSON(http.StatusCreated, data)
}

// Perbarui godoc
//
//	@Summary		Perbarui pemindaian xray
//	@Tags			PemindaianXray
//	@Accept			multipart/form-data
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string	true	"ID Pemindaian"
//	@Param			nama_image			formData	string	false	"Nama image container"
//	@Param			tanggal_pemindaian	formData	string	false	"Tanggal pemindaian (RFC3339)"
//	@Param			deskripsi			formData	string	false	"Deskripsi"
//	@Param			laporan				formData	string	false	"Laporan hasil pemindaian"
//	@Param			gambar				formData	file	false	"File gambar hasil scan"
//	@Success		200					{object}	models.PemindaianXray
//	@Failure		400					{object}	map[string]string
//	@Router			/api/v1/pemindaian-xray/{id} [put]
func (p *PengendaliPemindaianXray) Perbarui(c *gin.Context) {
	id := c.Param("id")

	namaImage := c.PostForm("nama_image")
	if namaImage == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama_image wajib diisi"})
		return
	}

	tglStr := c.DefaultPostForm("tanggal_pemindaian", time.Now().Format(time.RFC3339))
	tgl, err := time.Parse(time.RFC3339, tglStr)
	if err != nil {
		tgl = time.Now()
	}

	deskripsi := c.PostForm("deskripsi")
	laporan := c.PostForm("laporan")

	pathGambar := ""
	file, err := c.FormFile("gambar")
	if err == nil {
		ext := filepath.Ext(file.Filename)
		namaFile := uuid.New().String() + ext
		pathGambar = filepath.Join(p.Cfg.FolderGambar, namaFile)

		if err := c.SaveUploadedFile(file, pathGambar); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan gambar"})
			return
		}
	}

	data, err := p.Layanan.Perbarui(id, namaImage, deskripsi, pathGambar, laporan, tgl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Hapus godoc
//
//	@Summary		Hapus pemindaian xray
//	@Tags			PemindaianXray
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"ID Pemindaian"
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/api/v1/pemindaian-xray/{id} [delete]
func (p *PengendaliPemindaianXray) Hapus(c *gin.Context) {
	id := c.Param("id")

	ada, err := p.Layanan.AmbilBerdasarkanID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data tidak ditemukan"})
		return
	}

	if ada.PathGambar != "" {
		os.Remove(ada.PathGambar)
	}

	if err := p.Layanan.Hapus(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "data berhasil dihapus"})
}
