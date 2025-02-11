package data

import (
	"database/sql"
	"fmt"

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
	_, err := s.db.Query("INSERT INTO workers (Firstname, Secondname, Salary, Email, Post) values (?, ?, ?, ?, ?)",
	worker.Firstname, worker.Secondname, worker.Salary, worker.Email, worker.Post)
	return worker.UserID, err
}

func (s *Storage) Update(id uint64, worker *models.Worker) error {
	query := `
        UPDATE workers
        SET Firstname = ?, Secondname = ?, Salary = ?, Email = ?, Post = ?
        WHERE UserID = ?
    `
	result, err := s.db.Exec(query, worker.Firstname, worker.Secondname, worker.Salary, worker.Email, worker.Post, id)
    if err != nil {
        return fmt.Errorf("error updating worker: %w", err)
    }
	affected, err := result.RowsAffected()
	if err != nil{
		return fmt.Errorf("error getting affected: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("no worker found with id = %d", id)
	}

	return nil
}

func (s *Storage) Delete(id uint64) error {
	query := `
	DELETE FROM workers
    WHERE UserID = ?
	`

	result, err := s.db.Exec(query, id)
	if err != nil{
		return fmt.Errorf("error deleting worker id: %d, error: %w", id, err)
	}

	affected, err := result.RowsAffected()
	if err != nil{
		return fmt.Errorf("error get affected: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("worker not found")
	}

	return nil
}

func (s *Storage) Get(id uint64) (*models.Worker, error) {
	query := `
		SELECT FROM workers
		WHERE UserID = ?
	`

	result, err := s.db.Query(query, id)
	if err != nil{
		return nil, fmt.Errorf("error get, id: %d, error: %w", id, err)
	}

	var worker models.Worker
	for result.Next() {
		result.Scan(&worker.UserID, 
			&worker.Firstname, 
			&worker.Secondname,
			&worker.Salary,
			&worker.Email,
			&worker.Post)
	}

	return &worker, nil
}