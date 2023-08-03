package files

type AsrFiles struct {
	CurrentPage int    `json:"current_page"`
	Data        []File `json:"data"`
	TotalPages  int    `json:"total_pages"`
}
