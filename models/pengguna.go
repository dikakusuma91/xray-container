package models

import "time"

type Pengguna struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	KataSandi   string    `json:"-"`
	NamaLengkap string    `json:"nama_lengkap"`
	DibuatPada  time.Time `json:"dibuat_pada"`
}
