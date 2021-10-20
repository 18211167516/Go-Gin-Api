package config

type Captcha struct {
	KeyLen    int `json:"keyLen" toml:"key-len"`
	ImgWidth  int `json:"imgWidth" toml:"img-width"`
	ImgHeight int `json:"imgHeight" toml:"img-height"`
}
