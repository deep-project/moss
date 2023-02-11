package utils

import (
	"errors"
	"github.com/robfig/cron/v3"
)

// CheckCronExp 验证定时器表达式
func CheckCronExp(cronExp string) error {
	if _, err := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor).Parse(cronExp); err != nil {
		return errors.New("cron exp error" + err.Error())
	}
	return nil
}
