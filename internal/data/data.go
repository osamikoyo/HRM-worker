package data

import (
	"database/sql"
	"github.com/osamikoyo/hrm-worker/internal/config"
	"github.com/osamikoyo/hrm-worker/internal/data/models"
)

type Storage struct{
	db *sql.DB
}

func InitStorage(cfg *config.Config) (*Storage, error) {
	db, err := sql.Open("libsql", cfg.DatabaseURL)
	if err != nil {
	  return nil, err
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS workers (
    UserID INTEGER PRIMARY KEY AUTOINCREMENT,
    Firstname TEXT NOT NULL,
    Secondname TEXT NOT NULL,
    Salary INTEGER NOT NULL,
    Email TEXT NOT NULL UNIQUE,
    Post TEXT NOT NULL
);`

	_, err = db.Query(query)
	if err != nil{
		return nil, err
	}

	return &Storage{db}, err
}

func (s *Storage) Create(worker *models.Worker) (uint64, error) {
	_, err := s.db.Query("INSERT INTO workers (Firstname, Secondname, Salary, Email, Post) values (?1, ?2, ?3, ?4, ?5)",
	worker.Firstname, worker.Secondname, worker.Salary, worker.Email, worker.Post)
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