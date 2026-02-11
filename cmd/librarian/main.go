package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
	"google.golang.org/grpc"

	"Librarian/internal/app/librarian"
	"Librarian/internal/pkg/domain/presenter"
	"Librarian/internal/pkg/domain/usecase"
	"Librarian/internal/pkg/service/repository"
	"Librarian/internal/pkg/service/scheduler"
	"Librarian/internal/pkg/service/storage"
	"Librarian/internal/pkg/service/yandex"

	_ "modernc.org/sqlite"
)

// TODO сделать нормальную обработку ошибок
func main() {
	ctx := context.Background()

	//Получение яндекс клиента
	yandexClient := yandex.NewYandexDiskClient(os.Getenv("YANDEX_TOKEN"))
	fmt.Println("Яндекс клиент получен")

	db, err := storage.InitDB()
	if err != nil {
		fmt.Println("Подключиться к БД указателей с яндекс диска не удалось - ", err)
	}
	storageImpl := storage.NewStorage(db)

	presenterImpl := presenter.New()

	repositoryImpl := repository.NewRepository(storageImpl)

	usacase := usecase.NewUsacase(presenterImpl, repositoryImpl, yandexClient)

	// собираем сервер перед запуском
	server := librarian.NewServer(
		presenterImpl,
		usacase,
	)

	// собираем сервис фоновых задач по расписанию
	schedule := scheduler.NewScheduler(presenterImpl, usacase)

	schedule.DownloadSourceFile(ctx, 1*time.Minute) //TODO заменить на константу или глобальную переменную

	// gRPC сервис подключаем
	lis, err := net.Listen("tcp", os.Getenv("LIBRARIAN_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.ConnectionTimeout(10 * time.Second))
	lib.RegisterLibrarianServer(s, server)

	log.Printf("Librarion listening on %s", os.Getenv("LIBRARIAN_PORT"))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
