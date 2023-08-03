package files

import "time"

type File struct {
	FileId   string    `json:"file_id"`
	FileTime time.Time `json:"file_time"`
	FileType string    `json:"file_type"`
	Status   string    `json:"status"`
}
type Files struct {
	CurrentPage int    `json:"current_page"`
	Data        []File `json:"data"`
	TotalPages  int    `json:"total_pages"`
}
