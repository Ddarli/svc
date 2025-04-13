package repository

const (
	insertRecord = `INSERT INTO stored_files(
			id, user_id, file_name, description, file_size, mime_type, s3_key, hash, created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, NOW()
		)`
	selectRecord = `select id, user_id, file_name, description, file_size, mime_type, s3_key, hash 
		from stored_files 
		where user_id=$1`
	selectByID = `SELECT id, user_id, file_name, description, file_size, mime_type, s3_key, hash
		FROM stored_files
		WHERE id = $1`
)
