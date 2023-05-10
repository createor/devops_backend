package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type HttpsProtocol struct {
	IsUse   bool   `yaml:"enable"`
	Port    int16  `yaml:"port"`
	KeyPath string `yaml:"key"`
	PemPath string `yaml:"pem"`
}

type LogSetting struct {
	Level string `yaml:"level"`
}

type DBSetting struct {
	Name string `yaml:"dbname"`
}

// 配置
type Config struct {
	Port         int8          `yaml:"port"`       // 服务运行端口
	Secert       string        `yaml:"secret"`     //
	Expire       int8          `yaml:"expire"`     //
	DB           DBSetting     `yaml:"database"`   //
	InitPassword string        `yaml:"initpasswd"` //
	IgnorePath   []string      `yaml:"ignore"`     //
	Https        HttpsProtocol `yaml:"https"`      //
	Logger       LogSetting    `yaml:"logger"`     //
	FilePath     string        `yaml:"savepath"`   //
}

// Yaml2Struct yaml文件转换为结构体
//
// 参数:
//
//	filepath: string,yaml文件路径
//
// 返回:
//
//	Config: 配置结构体
//	error: 错误
func Yaml2Struct(filepath string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
