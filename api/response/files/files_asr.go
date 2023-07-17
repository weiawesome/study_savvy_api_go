package files

type AsrFiles struct {
	CurrentPage int `json:"current_page"`
	Data        []struct {
		FileId   string `json:"file_id"`
		FileTime string `json:"file_time"`
		FileType string `json:"file_type"`
		Status   string `json:"status"`
	} `json:"data"`
	TotalPages int `json:"total_pages"`
}
