package files

type OcrFiles struct {
	CurrentPage int    `json:"current_page"`
	Data        []File `json:"data"`
	TotalPages  int    `json:"total_pages"`
}
