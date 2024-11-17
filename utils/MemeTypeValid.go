package utils

func IsValidMimeType(mimeType string) bool {
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
