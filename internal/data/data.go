package data

import (
	"github.com/osamikoyo/hrm-worker/internal/data/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct{
	db *gorm.DB
}

func InitStorage(dsn string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil{
		return nil, err
	}
	err = db.AutoMigrate(&models.Worker{})

	return &Storage{db}, err
}

func (s *Storage) Create(worker *models.Worker) (uint64, error) {
	err := s.db.Create(worker).Error
	return worker.UserID, err
}

func (s *Storage) Update(id uint64, worker *models.Worker) error {
	result := s.db.Model(&models.Worker{}).Where(&models.Worker{
		UserID: id,
	}).Updates(worker)
	return result.Error
}

func (s *Storage) Delete(id uint64) error {
	result := s.db.Where(&models.Worker{
		UserID: id,
	}).Delete(&models.Worker{})
	return result.Error
}

func (s *Storage) Get(id uint64) (*models.Worker, error) {
	var worker models.Worker
	
	result := s.db.Where(&models.Worker{
		UserID: id,
	}).Find(&worker)
	return &worker, result.Error
}