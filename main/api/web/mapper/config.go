package mapper

import (
	"moss/api/web/dto"
	"moss/domain/config/service"
)

var Config = new(configMapper)

type configMapper struct {
}

// ConfigListToInfoList 把config指针集合转换成dto.ConfigInfo集合
func (c *configMapper) ConfigListToInfoList(maps []service.Config) (ret []dto.ConfigInfo) {
	for _, m := range maps {
		ret = append(ret, dto.ConfigInfo{ID: m.ConfigID(), Data: m})
	}
	return
}

func (c *configMapper) BodyToAdminInit(body []byte) (*dto.ConfigAdminInit, error) {
	var obj dto.ConfigAdminInit
	if err := BodyParser(body, &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}

func (c *configMapper) BodyToAdminLogin(body []byte) (*dto.ConfigAdminLogin, error) {
	var obj dto.ConfigAdminLogin
	if err := BodyParser(body, &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}

func (c *configMapper) BodyToAdminPost(body []byte) (*dto.ConfigAdminPost, error) {
	var obj dto.ConfigAdminPost
	if err := BodyParser(body, &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}
