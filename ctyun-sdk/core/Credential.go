package core

type Credential struct {
	AccessKey string
	SecretKey string
}

func NewCredentials(accessKey, secretKey string) *Credential {
	return &Credential{accessKey, secretKey}
}
