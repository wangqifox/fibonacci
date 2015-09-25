package main

import (
	"github.com/Unknwon/log"
	"github.com/Unknwon/macaron"
	"github.com/wangqifox/fibonacci/modules/setting"
	"github.com/wangqifox/fibonacci/routers/v1"
	"net/http"
)

const APP_VER = "0.0.1.0912"

func main() {
	log.Info("App Version: %s", APP_VER)
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	// m.Use(middleware.Contexter())

	m.Get("/", func() string {
		return "Hello"
	})
	m.Get("/fibonacci", v1.Fibonacci)

	log.Info("PORT: %s", setting.HTTPPort)
	http.ListenAndServe(":"+setting.HTTPPort, m)
	// m.Run(":" + setting.HTTPPort)
}
