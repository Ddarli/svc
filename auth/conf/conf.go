package conf

import "auth/internal/service"

type Configuration struct {
	Server struct {
		Port string
	}
	Routes struct {
		Register     string
		Authenticate string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	BlockchainClient struct {
		Port string
	}

	service.ServiceConf
}
