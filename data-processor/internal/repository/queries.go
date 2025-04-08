package repository

const (
	insertRecord = `insert into stored_files(id, user_id, file_name, file_size, mime_type, s3_key, hash, created_at) values 
    	($1, $2, $3, $4, $5, $6, $7, now())`
	selectRecord = `select (id, user_id, file_name, file_size, mime_type, s3_key, hash, created_at) from stored_files where user_id=$1`
)
