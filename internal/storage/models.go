package storage

import "gorm.io/gorm"

type Station struct {
	gorm.Model
	Name string
}
