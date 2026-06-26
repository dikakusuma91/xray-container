package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"xray-post-test/config"
	"xray-post-test/controllers"
	"xray-post-test/middleware"
	"xray-post-test/repository"
	"xray-post-test/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DaftarkanRute(router *gin.Engine, db *sql.DB, cfg *config.Konfigurasi) {
	// Repository
	repoPengguna := repository.BuatRepositoriPengguna(db)
	repoPemindaian := repository.BuatRepositoriPemindaianXray(db)

	// Services
	layananAuth := services.BuatLayananAuth(repoPengguna, cfg)
	layananPemindaian := services.BuatLayananPemindaianXray(repoPemindaian)

	// Controllers
	pengendaliAuth := controllers.BuatPengendaliAuth(layananAuth)
	pengendaliPemindaian := controllers.BuatPengendaliPemindaianXray(layananPemindaian, cfg)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth routes (publik)
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/daftar", pengendaliAuth.Daftar)
		auth.POST("/masuk", pengendaliAuth.Masuk)
	}

	// Protected routes (JWT)
	protected := router.Group("/api/v1/pemindaian-xray")
	protected.Use(middleware.PerantaraJWT(cfg))
	{
		protected.GET("", pengendaliPemindaian.AmbilSemua)
		protected.GET("/:id", pengendaliPemindaian.AmbilBerdasarkanID)
		protected.POST("", pengendaliPemindaian.BuatBaru)
		protected.PUT("/:id", pengendaliPemindaian.Perbarui)
		protected.DELETE("/:id", pengendaliPemindaian.Hapus)
	}
}
