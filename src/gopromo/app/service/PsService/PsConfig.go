package PsService

type imageParams struct {
	Host  string
	Name  string
	Field []string
}

var psConfig = map[string]imageParams{
	"jdbutterfly": imageParams{"115.238.138.134:8900", "jdbutterfly", []string{"image_url"}},
}

func getConfigValue(key string) (imageParams, bool) {
	value, ok := psConfig[key]
	return value, ok
}
