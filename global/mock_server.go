package global

var (
	MockServerDict map[string]map[string]MockServerItemConfig
)

type MockServerItemConfig struct {
	Uri      string
	Method   string
	Response string
}
