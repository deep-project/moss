package controller

import (
	"github.com/gofiber/fiber/v2"
	"moss/api/web/dto"
	"moss/api/web/mapper"
	"moss/application/service"
	"moss/infrastructure/general/constant"
	"sync"
	"time"
)

func AdminExists(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.AdminExists(), nil))
}

func AdminCreate(ctx *fiber.Ctx) error {
	if time.Since(constant.AppStartTime).Minutes() > 10 {
		return ctx.JSON(mapper.MessageFail("please restart the application to create an administrator within 10 minutes"))
	}
	obj, err := mapper.Config.BodyToAdminInit(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.AdminCreate(obj.Username, obj.Password)))
}

var loginLock sync.Mutex

func AdminLogin(ctx *fiber.Ctx) error {
	loginLock.Lock()
	defer loginLock.Unlock()
	time.Sleep(1200 * time.Millisecond)
	obj, err := mapper.Config.BodyToAdminLogin(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.AdminLogin(obj.Username, obj.Password, obj.Captcha, obj.CaptchaID)))
}

func AdminCaptcha(ctx *fiber.Ctx) error {
	loginLock.Lock()
	defer loginLock.Unlock()
	time.Sleep(1200 * time.Millisecond)
	return ctx.JSON(mapper.MessageResultData(dto.NewCaptcha(service.AdminCaptcha()), nil))
}
