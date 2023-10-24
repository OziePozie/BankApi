package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	const op = "storage.postgres.NewStorage" // Имя текущей функции для логов и ошибок

	db, err := sql.Open("postgres",
		"user=postgres password=123 host=localhost dbname=bankdb sslmode=disable") // Подключаемся к БД
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Создаем таблицу, если ее еще нет

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Get() *sql.DB {
	return s.db
}
