package entity

import "strconv"

type TLS struct {
	Enable         bool   `json:"enable"`
	KeyPEM         string `json:"key_pem"`
	CertPEM        string `json:"cert_pem"`
	Port           int    `json:"port"`
	ForceHTTPS     bool   `json:"force_https"`     // 强制跳转到 https
	RedirectStatus int    `json:"redirect_status"` // 重定向状态码
}

func NewTLS() *TLS {
	return &TLS{Enable: false, Port: 443, RedirectStatus: 302}
}

func (*TLS) ConfigID() string {
	return "tls"
}

func (t *TLS) ListenPort() int {
	if t.Port > 0 {
		return t.Port
	}
	return 443
}

func (t *TLS) GetRedirectStatus() int {
	if t.RedirectStatus == 301 {
		return 301
	}
	return 302
}

func (t *TLS) ListenAddr() string {
	return ":" + strconv.Itoa(t.ListenPort())
}
