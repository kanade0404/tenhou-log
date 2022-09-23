BEGIN;
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE log_file_packages (
    file_name VARCHAR(255) PRIMARY KEY,
    size INTEGER NOT NULL
);
CREATE TABLE log_files (
    id VARCHAR PRIMARY KEY DEFAULT gen_random_uuid()
);
COMMIT;
