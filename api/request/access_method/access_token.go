package access_method

type AccessToken struct {
	AesKey      string `json:"aes_key"`
	AccessToken string `json:"access_token"`
}
