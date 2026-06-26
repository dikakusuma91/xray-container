CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE pengguna (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email         VARCHAR(255) UNIQUE NOT NULL,
    kata_sandi    VARCHAR(255) NOT NULL,
    nama_lengkap  VARCHAR(255) DEFAULT '',
    dibuat_pada   TIMESTAMP NOT NULL DEFAULT NOW()
);
