package handlers

import (
	"gorm.io/gorm"
	"time"
)

type API struct {
	DB *gorm.DB
	OrderMap map[uint]time.Time
}