CREATE TABLE pemindaian_xray (
    id                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    id_pengguna         UUID NOT NULL REFERENCES pengguna(id),
    nama_image          VARCHAR(255) NOT NULL,
    tanggal_pemindaian  TIMESTAMP NOT NULL DEFAULT NOW(),
    deskripsi           TEXT DEFAULT '',
    path_gambar         VARCHAR(500) DEFAULT '',
    prioritas           VARCHAR(10) NOT NULL CHECK (prioritas IN ('rendah', 'sedang', 'tinggi')),
    laporan             TEXT DEFAULT '',
    dibuat_pada         TIMESTAMP NOT NULL DEFAULT NOW(),
    diperbarui_pada     TIMESTAMP NOT NULL DEFAULT NOW()
);
