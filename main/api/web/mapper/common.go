package mapper

import (
	"encoding/json"
	"moss/api/web/dto"
	"moss/domain/config"
	"moss/domain/core/aggregate"
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/utils"
)

func MessageResult(err error) *dto.MessageResult {
	if err == nil {
		return &dto.MessageResult{Success: true}
	}
	return &dto.MessageResult{Message: err.Error()}
}

func MessageResultData(data any, err error) *dto.MessageResult {
	if err == nil {
		return &dto.MessageResult{Success: true, Data: data}
	}
	return &dto.MessageResult{Message: err.Error()}
}

func MessageFail(msg string) *dto.MessageResult {
	return &dto.MessageResult{Message: msg}
}

func MessageSuccess(msg string) *dto.MessageResult {
	return &dto.MessageResult{Success: true, Message: msg}
}

func BodyParser(body []byte, ptr any) error {
	return json.Unmarshal(body, ptr)
}

func BodyToContext(body []byte) (ctx context.Context, err error) {
	if len(body) == 0 {
		return
	}
	err = BodyParser(body, &ctx)
	ctx.FastOffset = config.Config.More.FastOffsetMinPage > 0 && ctx.Page > config.Config.More.FastOffsetMinPage // 加速分页查询
	if ctx.Limit == 0 {
		ctx.Limit = 20 // 限制调取数量
	}
	ctx.Order = utils.ValidateOrderInput(ctx.Order)
	return
}

func BodyToWhere(body []byte) (res context.Where, err error) {
	if len(body) == 0 {
		return
	}
	err = BodyParser(body, &res)
	return
}

type curdModel interface {
	entity.Article | entity.Category | entity.Tag | entity.Link | aggregate.ArticlePost | entity.Store
}

func BodyToStrSet(body []byte) (res []string, err error) {
	err = BodyParser(body, &res)
	return
}

func BodyToIntSet(body []byte) (res []int, err error) {
	err = BodyParser(body, &res)
	return
}
