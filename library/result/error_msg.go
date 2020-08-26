package result

import "sync"

const (
	defaultZhCnErrMsg    = "抱歉，发生了一个错误"
	defaultZhTwErrMsg    = "抱歉，发生了一个错误"
	defaultEnglishErrMsg = "sorry, system have a error"
)

var (
	errorsMap sync.Map // key: string value: Error
)

type errorMsg struct {
	code  int
	zhMsg string
	twMsg string
	enMsg string
}

func FindErrorMessage(key string, language string) (string, int) {
	err, ok := errorsMap.Load(key)
	errV, _ := err.(errorMsg)

	switch language {
	case "zh":
		if !ok || errV.zhMsg == "" {
			return defaultZhCnErrMsg, 0
		}

		return errV.zhMsg, errV.code
	case "tw":
		if !ok || errV.twMsg == "" {
			return defaultZhTwErrMsg, 0
		}

		return errV.twMsg, errV.code
	case "en":
		if !ok || errV.enMsg == "" {
			return defaultEnglishErrMsg, 0
		}

		return errV.enMsg, errV.code
	default:
		return defaultZhCnErrMsg, 0
	}
}

func InsertError(lang string, code int, key string, msg string) {
	var err errorMsg
	value, ok := errorsMap.Load(key)

	if !ok {
		err = errorMsg{}
	} else {
		err = value.(errorMsg)
	}

	switch lang {
	case "zh":
		err.zhMsg = msg
	case "tw":
		err.twMsg = msg
	case "en":
		err.enMsg = msg
	}

	err.code = code
	errorsMap.Store(key, err)
}
