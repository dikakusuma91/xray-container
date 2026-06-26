package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Konfigurasi struct {
	Port         string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	AppEnv       string
	JwtRahasia   string
	FolderGambar string
}

func MuatKonfigurasi() *Konfigurasi {
	godotenv.Load()

	return &Konfigurasi{
		Port:         getEnv("PORT", "8181"),
		DBHost:       getEnv("DB_HOST", "127.0.0.1"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "admin123"),
		DBName:       getEnv("DB_NAME", "go_gin_crud"),
		AppEnv:       getEnv("APP_ENV", "development"),
		JwtRahasia:   getEnv("JWT_RAHASIA", "rahasia_default"),
		FolderGambar: getEnv("FOLDER_GAMBAR", "repo-xray"),
	}
}

func (k *Konfigurasi) DSN() string {
	return "host=" + k.DBHost +
		" port=" + k.DBPort +
		" user=" + k.DBUser +
		" password=" + k.DBPassword +
		" dbname=" + k.DBName +
		" sslmode=disable"
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
