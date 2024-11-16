package models

type ArchiveFile struct {
	Filename     string  `json:"filename"`
	Archive_size float32 `json:"archive_size"`
	Totalsize    float32 `json:"total_size"`
	Total_files  float32 `json:"total_files"`
	Files        []File  `json:"files"`
}

type File struct {
	File_path string  `json:"file_path"`
	Size      float64 `json:"size"`
	Mimetype  string  `json:"mimetype"`
}
