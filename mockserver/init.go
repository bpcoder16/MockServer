package mockserver

import (
	"MockServer/global"
	"github.com/bpcoder16/Mario/utils"
)

func init() {
	var mockServerConfigMap map[string][]global.MockServerItemConfig
	utils.SetConfigWithFile("./config/mock_server.toml", &mockServerConfigMap)
	global.MockServerDict = make(map[string]map[string]global.MockServerItemConfig)
	for port, mockServerItemConfigList := range mockServerConfigMap {
		global.MockServerDict[port] = make(map[string]global.MockServerItemConfig)
		for _, item := range mockServerItemConfigList {
			global.MockServerDict[port][item.Uri] = global.MockServerItemConfig{
				Method:   item.Method,
				Response: item.Response,
			}
		}
	}
}
