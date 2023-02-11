package entity

import "github.com/duke-git/lancet/v2/random"

type API struct {
	Enable    bool   `json:"enable"`
	SecretKey string `json:"secret_key"`
}

func NewAPI() *API {
	return &API{Enable: false, SecretKey: random.RandString(20)}
}

func (*API) ConfigID() string {
	return "api"
}
