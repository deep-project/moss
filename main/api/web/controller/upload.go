package controller

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"moss/api/web/mapper"
	"moss/infrastructure/persistent/storage"
	"moss/infrastructure/support/upload"
)

func UploadInit(ctx *fiber.Ctx) error {
	if err := upload.Init(); err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(nil))
}

func Upload(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	var urls []string
	for _, header := range form.File {
		for _, v := range header {
			res, err := uploadFunc(v)
			if err != nil {
				return ctx.JSON(mapper.MessageResult(err))
			}
			urls = append(urls, res.URL)
		}
	}
	return ctx.JSON(mapper.MessageResultData(urls, nil))
}

func uploadFunc(h *multipart.FileHeader) (res *upload.Result, err error) {
	f, err := h.Open()
	if err != nil {
		return
	}
	defer f.Close()
	return upload.Upload(h.Filename, "", storage.NewSetValue(f))
}
