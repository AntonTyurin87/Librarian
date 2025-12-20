package sqlite

import (
	"Librarian/internal/yandex"
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

// Storage - Структура для хранилища
type Storage struct {
	DB *sql.DB
}

// NewStorage - онструктор для хранилица
func NewStorage(db *sql.DB) Storage {
	return Storage{DB: db}
}

var LibrarianStorage Storage

// InitDB - создаёт подключение к базе данных
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", yandex.CurrentDir)
	if err != nil {
		fmt.Println("Нет подключения к базе данных", err)
		return nil, err
	}

	return db, nil
}
