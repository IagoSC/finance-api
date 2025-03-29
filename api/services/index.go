package services

import "gorm.io/gorm"

type service struct {
	DB *gorm.DB
}
