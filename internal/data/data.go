package data

import (
	"github.com/osamikoyo/hrm-worker/internal/data/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct{
	*gorm.DB
}

func InitStorage(dsn string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil{
		return nil, err
	}
	err = db.AutoMigrate(&models.Worker{})

	return &Storage{db}, err
}

func (s *Storage) Create(worker *models.Worker) error {
	
}