package entity

type TLS struct {
	Enable     bool   `json:"enable"`
	KeyPEM     string `json:"key_pem"`
	CertPEM    string `json:"cert_pem"`
	ListenAddr string `json:"listen_addr"` // 监听地址 默认 :443
}

func NewTLS() *TLS {
	return &TLS{Enable: false, ListenAddr: ":443"}
}

func (*TLS) ConfigID() string {
	return "tls"
}
