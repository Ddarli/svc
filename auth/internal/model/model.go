package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	User struct {
		ID        uuid.UUID
		Username  string
		Password  string
		Email     string
		Phone     string
		Role      int
		CreatedAt time.Time
	}
)
