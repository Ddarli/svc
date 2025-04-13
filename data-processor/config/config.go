package config

type Config struct {
	Server struct {
		Port string
		Mode string
	}
	Storage struct {
		Bucket   string
		SecretID string
		Secret   string
	}
	Database struct {
		Host     string
		User     string
		Password string
		Port     string
		Name     string
	}
}
