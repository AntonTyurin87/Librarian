package yandex

import (
	"fmt"
	"log"
)

var CurrentDir string

func GetLibrarianDB(client *YandexDiskClient) {
	// Путь к файлу на Яндекс Диске
	remoteFilePath := "/ReconCom/librarian.sqlite"

	// Получаем текущую директорию проекта
	var err error
	CurrentDir, err = GetCurrentDir()
	if err != nil {
		log.Fatalf("❌ Ошибка получения текущей директории: %v", err)
	}

	fmt.Printf("📁 Текущая директория проекта: %s\n", CurrentDir)
	fmt.Printf("📁 Файл на Яндекс Диске: %s\n", remoteFilePath)

	// Скачиваем файл
	fmt.Println("\n🔄 Начинаем скачивание...")
	err = client.DownloadFile(remoteFilePath, CurrentDir)
	if err != nil {
		log.Fatalf("❌ Ошибка скачивания файла: %v", err)
	}

	fmt.Println("✅ Файл успешно скачан в директорию проекта!")
}
