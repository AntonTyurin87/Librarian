package main

import (
	"Librarian/internal/grpc/librarian"
	"Librarian/internal/yandex"
	"Librarian/internal/yandex/sqlite"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
)

func main() {
	// Получение яндекс клиента
	yandex.YandexClient = yandex.NewYandexDiskClient(os.Getenv("YANDEX_TOKEN"))
	fmt.Println("Яндекс клиент получен")

	// Получи файл базы данных яндекс диска
	yandex.GetLibrarianDB(yandex.YandexClient)
	fmt.Println("Яндекс клиент получен")

	// Подключились к БД указателей на источники яндекс диска
	db, err := sqlite.InitDB()
	if err != nil {
		fmt.Println("Подключиться к БД указателей с яндекс диска не удалось - ", err)
	}
	sqlite.LibrarianStorage = sqlite.NewStorage(db)

	// gRPC сервис подключаем
	librarian.InitGRPC()
}
