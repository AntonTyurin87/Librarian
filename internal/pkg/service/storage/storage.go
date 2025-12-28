package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Storage - Структура для хранилища
type Storage struct {
	DB *sql.DB
}

//	func (s Storage) Exec(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
//		//TODO implement me
//		panic("implement me")
//	}
func (s Storage) Query(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	return s.DB.QueryContext(ctx, sql, args...)
}

//
//func (s Storage) QueryRow(ctx context.Context, sql string, args ...interface{}) *sql.Row {
//	//TODO implement me
//	panic("implement me")
//}

// NewStorage - онструктор для хранилица
func NewStorage(db *sql.DB) Storage {
	return Storage{DB: db}
}

// InitDB - создаёт подключение к базе данных
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "RC_librarian.db") //TODO заменить на переменную или константу адрес БД
	if err != nil {
		fmt.Println("Нет подключения к базе данных", err)
		return nil, err
	}

	return db, nil
}
