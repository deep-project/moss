package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
	"moss/domain/config"
	"regexp"
)

func init() {
	initMinifyCodeHandle()
}

var minifyCodeHandle *minify.M

func MinifyCode(ctx *fiber.Ctx) error {
	next := ctx.Next()
	if !config.Config.Router.MinifyCode || next != nil || minifyCodeHandle == nil {
		return next
	}
	body := ctx.Response().Body()
	mediaType := string(ctx.Response().Header.ContentType())
	if len(body) == 0 {
		return next
	}
	body, err := minifyCodeHandle.Bytes(mediaType, body)
	if err != nil || len(body) == 0 {
		return next
	}
	ctx.Context().SetBody(body)
	return next
}

func initMinifyCodeHandle() {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.Add("text/html", &html.Minifier{
		KeepDefaultAttrVals:     true, // 保留标签的默认属性
		KeepDocumentTags:        true, // 删除多余的标记，html、head和body标记除外
		KeepQuotes:              true, // 保留引号
		KeepEndTags:             true, // 保留 end tags
		KeepConditionalComments: true, // 保留条件注释  <!--[if 。。。<![endif]-->
	})
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	minifyCodeHandle = m
}
