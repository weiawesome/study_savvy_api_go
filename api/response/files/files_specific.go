package files

type SpecificFile struct {
	Content   string   `json:"content"`
	Details   []string `json:"details"`
	Prompt    string   `json:"prompt"`
	Summarize string   `json:"summarize"`
}
