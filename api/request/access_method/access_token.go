package access_method

type AccessToken struct {
	AesKey      string `json:"AES_key"`
	AccessToken string `json:"access_token"`
}
