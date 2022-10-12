CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE compressed_log_files (
                                      name VARCHAR(255) PRIMARY KEY,
                                      size INTEGER NOT NULL
);
CREATE TABLE log_files (
    id VARCHAR PRIMARY KEY DEFAULT gen_random_uuid()
);
