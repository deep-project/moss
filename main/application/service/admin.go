package service

import (
	"errors"
	"moss/domain/config"
	"moss/domain/config/service"
	"moss/infrastructure/support/captcha"
	"moss/infrastructure/utils/timex"
)

// AdminExists 判断管理员是否存在
func AdminExists() bool {
	return config.Config.Admin.Username != ""
}

// AdminCreate 创建管理员
func AdminCreate(username, password string) error {
	if AdminExists() {
		return errors.New("administrator already exists")
	}
	err := config.Config.Admin.InitAdministrator(username, password)
	if err != nil {
		return err
	}
	return service.Push(config.Config.Admin)
}

func AdminLogin(username, password, captchaAnswer, captchaID string) (token string, err error) {
	if err = captcha.Client.Verify(captchaID, captchaAnswer); err != nil {
		return
	}
	captcha.Client.Delete(captchaID) // 验证码成功后，清除旧的验证码，防止重复使用
	return config.Config.Admin.Login(username, password)
}

func AdminCaptcha() (bs64 string, id string) {
	return captcha.Client.StringSimple(4).Base64()
}

// AdminUpdate 更新管理员配置
func AdminUpdate(username, password string, loginExpire timex.Duration) error {
	if err := config.Config.Admin.Update(username, password, loginExpire); err != nil {
		return err
	}
	config.Config.Admin.ResetJwtKey() // 重置 jwtKey主动使所有已登录失效
	return service.Push(config.Config.Admin)
}

// AdminUsernameUpdate 更新管理员用户名
func AdminUsernameUpdate(username string) error {
	if err := config.Config.Admin.UpdateUsername(username); err != nil {
		return err
	}
	config.Config.Admin.ResetJwtKey()
	return service.Push(config.Config.Admin)
}

// AdminPasswordUpdate 更新管理员密码
func AdminPasswordUpdate(password string) error {
	if err := config.Config.Admin.UpdatePassword(password); err != nil {
		return err
	}
	config.Config.Admin.ResetJwtKey()
	return service.Push(config.Config.Admin)
}

// AdminPathUpdate 更新管理路径
func AdminPathUpdate(path string) error {
	if err := config.Config.Router.UpdateAdminPath(path); err != nil {
		return err
	}
	return service.Push(config.Config.Router)
}
