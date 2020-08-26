package initialize

import (
	"basic_app/assets"
	"basic_app/conf"
	"errors"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

var (
	YamlFileNotFoundErr = errors.New("配置文件不存在")
	YamlFormatErr       = errors.New("解析config.yml失败，检查yaml格式是否有误")
	YamlMapStructureErr = errors.New("map转换struct失败，请联系开发者")
)

func InitConfig() (err error) {
	var data = make(map[string]interface{})

	configData, err := assets.Asset("conf/config.yaml")
	if err != nil {
		return YamlFileNotFoundErr
	}

	if err = yaml.Unmarshal(configData, data); err != nil {
		return YamlFormatErr
	}

	if err = mapstructure.Decode(data, &conf.Conf); err != nil {
		return YamlMapStructureErr
	}

	return
}
