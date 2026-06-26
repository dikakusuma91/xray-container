package main

import (
	"github.com/gin-gonic/gin"
	"xray-post-test/config"
	"xray-post-test/database"
	"xray-post-test/routes"

	_ "xray-post-test/docs"
)

// @title			X-Ray Container Scan API
// @version		1.0
// @description	API untuk mengelola hasil pemindaian container X-Ray
// @host			localhost:8181
// @BasePath		/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	cfg := config.MuatKonfigurasi()
	db := database.SambungKePostgres(cfg)
	defer db.Close()

	router := gin.Default()
	routes.DaftarkanRute(router, db, cfg)

	router.Run(":" + cfg.Port)
}
