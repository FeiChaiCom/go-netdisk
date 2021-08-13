package misc

import "encoding/json"

func PrettyJson(v interface{}) string {
	beauty, _ := json.MarshalIndent(v, "", "  ")
	return string(beauty)
}
