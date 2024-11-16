package utils

import "mime/multipart"

func IsValidMimeType(file *multipart.FileHeader) bool {
	// Получаем MIME тип файла
	mimeType := file.Header.Get("Content-Type")
	// Проверяем, является ли MIME тип допустимым
	validMimeTypes := []string{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document", // docx
		"application/pdf", // pdf
	}

	for _, validType := range validMimeTypes {
		if mimeType == validType {
			return true
		}
	}
	return false
}
