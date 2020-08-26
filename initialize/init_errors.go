package initialize

import (
	"basic_app/assets"
	"basic_app/library/result"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func InitErrorsMap() (err error) {
	var (
		content = make(map[string]string)
		data    []byte
	)

	loop := map[string]string{
		"zh": "locales/errors_zh.json",
		"tw": "locales/errors_tw.json",
		"en": "locales/errors_en.json",
	}

	for lang, filepath := range loop {

		if data, err = assets.Asset(filepath); err != nil {
			return err
		}

		if err = json.Unmarshal(data, &content); err != nil {
			return
		}

		writeToErrorsMap(lang, content)
		content = make(map[string]string)
		data = nil
	}

	return
}

func writeToErrorsMap(lang string, data map[string]string) {
	for key, value := range data {
		codestr := strings.SplitN(key, "_", 2)[0]
		code, _ := strconv.ParseUint(codestr, 10, 64)

		fmt.Println(code, int(code))

		result.InsertError(lang, int(code), key, value)
	}
}
