package main

import (
	"moss/api/web/router"
	_ "moss/startup"
)

func main() {
	err := router.New().Run()
	if err != nil {
		panic(err)
	}
}
