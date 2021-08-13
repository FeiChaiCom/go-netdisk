package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-netdisk/pkg/utils/codec"
	"os"
	"path/filepath"
	"strings"
)

type APIRespone struct {
	Result  bool        `json:"result"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func InternalError(code int, message string) error {

	_, err := createRespone(code, message, make(map[string]interface{}))

	return err
}

func InternalErrorEx(code int, message string) ([]byte, error) {
	return createResponeEx(code, message, make(map[string]interface{}))
}

func GetRespWithoutData(code int, message string) string {

	ret, _ := createRespone(code, message, make(map[string]interface{}))

	return ret
}

func GetRespWithoutDataEx(code int, message string) []byte {

	ret, _ := createResponeEx(code, message, make(map[string]interface{}))
	return ret
}

// deprecated: best to use GetResponse instead.
func GetRespone(code int, message string, data interface{}) string {

	ret, _ := createRespone(code, message, data)

	return ret
}

func createRespone(code int, message string, data interface{}) (string, error) {

	b, err := createResponeEx(code, message, data)

	return string(b), err
}

func createResponeEx(code int, message string, data interface{}) ([]byte, error) {
	bResult := false
	if 0 == code {
		bResult = true
	} else {
		appName := os.Args[0]
		szArr := strings.Split(appName, "/")
		if len(szArr) >= 2 {
			appName = szArr[1]
		}
		message = "(" + appName + "):" + message
	}

	resp := APIRespone{bResult, code, message, data}
	b, err := json.Marshal(resp)
	if err != nil {
		return []byte(""), err
	}

	return b, errors.New(string(b))
}

func GetResponse(code int, message string, data interface{}) ([]byte, error) {
	return createResponseEx(code, message, data, nil)
}

func GetResponseEx(code int, message string, data interface{}, extra map[string]interface{}) ([]byte, error) {
	return createResponseEx(code, message, data, extra)
}

func createResponseEx(code int, message string, data interface{}, extra map[string]interface{}) (r []byte, err error) {
	result := code == 0
	if !result {
		appName := filepath.Base(os.Args[0])
		message = fmt.Sprintf("(%s):%s", appName, message)
	}

	resp := APIRespone{result, code, message, data}
	if err = codec.EncJson(resp, &r); err != nil {
		return
	}

	return addExtraField(r, extra)
}

func addExtraField(s []byte, extra map[string]interface{}) (r []byte, err error) {
	if extra == nil {
		return s, nil
	}

	var jsn map[string]interface{}
	if err = codec.DecJson(s, &jsn); err != nil {
		return
	}
	for k, v := range extra {
		if _, ok := jsn[k]; !ok {
			jsn[k] = v
		}
	}
	err = codec.EncJson(jsn, &r)
	return
}
