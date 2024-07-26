package db

import (
	"sync"

	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Init() {
	once.Do(func() {

	})
}
