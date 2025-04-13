create table stored_files
(
    id             UUID PRIMARY KEY,
    user_id        UUID NOT NULL,
    file_name      TEXT NOT NULL,
    description    TEXT NOT NULL,
    status         TEXT,
    file_extension TEXT,
    file_size      BIGINT,
    mime_type      TEXT,
    s3_key         TEXT NOT NULL,
    hash           TEXT,
    created_at     TIMESTAMP DEFAULT now()
);