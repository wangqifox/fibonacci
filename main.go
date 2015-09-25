package main

import (
	// "flag"
	// "fmt"
	"github.com/Unknwon/log"
	"github.com/Unknwon/macaron"
	// "github.com/wangqifox/fibonacci/modules/setting"
	"github.com/wangqifox/fibonacci/routers/v1"
	"net/http"
	"os"
)

const APP_VER = "0.0.1.0912"

// var port = flag.Int("port", 4000, "Server listen port")

const (
	HostVar = "VCAP_APP_HOST"
	PortVar = "VCAP_APP_PORT"
)

func main() {
	// flag.Parse()
	log.Info("App Version: %s", APP_VER)
	var port string
	if port = os.Getenv(PortVar); port == "" {
		port = "8080"
	}
	log.Info("PORT: %s", port)

	m := macaron.Classic()
	m.Use(macaron.Renderer())
	// m.Use(middleware.Contexter())

	m.Get("/", func() string {
		return "Hello"
	})
	m.Get("/fibonacci", v1.Fibonacci)

	// log.Info("PORT: %s", setting.HTTPPort)
	// _ = setting.HTTPPort

	http.ListenAndServe(":"+port, m)

	// http.ListenAndServe(fmt.Sprintf(":%d", *port), m)
	// http.ListenAndServe(":"+setting.HTTPPort, m)
	// m.Run(":" + setting.HTTPPort)
}
