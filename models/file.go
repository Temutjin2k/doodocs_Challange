package models

type ArchiveFile struct {
	Filename     string  `json:"filename"`
	Archive_size float64 `json:"archive_size"`
	Totalsize    float64 `json:"total_size"`
	Total_files  int     `json:"total_files"`
	Files        []File  `json:"files"`
}

type File struct {
	File_path string  `json:"file_path"`
	Size      float64 `json:"size"`
	Mimetype  string  `json:"mimetype"`
}
