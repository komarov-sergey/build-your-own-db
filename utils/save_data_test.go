package utils

import (
	"os"
	"testing"
)

// TestSaveData2 тестирует функцию SaveData2
func TestSaveData2(t *testing.T) {
	// Создаем временный файл
	tmpFile, err := os.CreateTemp("", "test-*.tmp")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Удалим временный файл после теста

	// Данные для записи
	data := []byte("Hello, World!")

	// Вызываем SaveData2
	err = SaveData2(tmpFile.Name(), data)
	if err != nil {
		t.Fatalf("SaveData2 failed: %v", err)
	}

	// Проверяем, что данные были записаны корректно
	writtenData, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read written data: %v", err)
	}
	if string(writtenData) != string(data) {
		t.Errorf("Data mismatch: expected %s, got %s", data, writtenData)
	}

	// Проверяем, что временный файл был переименован в исходный путь
	if _, err := os.Stat(tmpFile.Name()); os.IsNotExist(err) {
		t.Errorf("File was not renamed to %s", tmpFile.Name())
	}
}
