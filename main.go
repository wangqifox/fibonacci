package main

import (
	"github.com/Unknwon/log"
	"github.com/Unknwon/macaron"
	"github.com/wangqifox/fibonacci/modules/middleware"
	"github.com/wangqifox/fibonacci/modules/setting"
	"github.com/wangqifox/fibonacci/routers/v1"
)

const APP_VER = "0.0.1.0912"

func main() {
	log.Info("App Version: %s", APP_VER)
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(middleware.Contexter())

	m.Get("/fibonacci", v1.Fibonacci)

	m.Run(setting.HTTPPort)
}
