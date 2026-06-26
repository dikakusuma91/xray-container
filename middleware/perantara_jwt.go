package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"xray-post-test/config"
	"xray-post-test/services"
)

func PerantaraJWT(cfg *config.Konfigurasi) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token tidak ditemukan"})
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		if tokenStr == header {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "format token salah"})
			return
		}

		klaim := &services.KlaimJWT{}
		token, err := jwt.ParseWithClaims(tokenStr, klaim, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JwtRahasia), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token tidak valid"})
			return
		}

		c.Set("id_pengguna", klaim.IDPengguna)
		c.Set("email", klaim.Email)
		c.Set("nama_lengkap", klaim.NamaLengkap)
		c.Next()
	}
}
