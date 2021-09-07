package main

import (
	"MockServer/global"
	_ "MockServer/mockserver"
	v1 "MockServer/mockserver/v1"
	"github.com/bpcoder16/Mario"
	"github.com/bpcoder16/Mario/core"
	"github.com/bpcoder16/Mario/mario"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	configList := make([]core.HttpServerConfig, 0, len(global.MockServerDict))
	for port, mockServerItemMap := range global.MockServerDict {
		indexHandler := v1.JsonApiPort(port)
		configList = append(configList, core.HttpServerConfig{
			Port: port,
			Manager: func(r *gin.Engine) {
				for uri, item := range mockServerItemMap {
					switch strings.ToUpper(item.Method) {
					case "GET":
						r.GET(uri, indexHandler)
					case "POST":
						r.POST(uri, indexHandler)
					}
				}
			},
		})
	}
	mario.ZapSugaredLogger.Info(configList)

	Mario.RunMultiHttpServer(configList)
}
