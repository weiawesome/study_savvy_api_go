package information

type EditPassword struct {
	NewPwd      string `json:"new_pwd"`
	OriginalPwd string `json:"original_pwd"`
}
