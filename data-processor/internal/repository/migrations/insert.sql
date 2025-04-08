create table stored_files
(
    id             UUID PRIMARY KEY,
    user_id        UUID REFERENCES users (id),
    file_name      TEXT NOT NULL,
    file_extension TEXT,
    file_size      BIGINT,
    mime_type      TEXT,
    s3_key         TEXT NOT NULL,
    hash           TEXT,
    created_at     TIMESTAMP DEFAULT now()
);