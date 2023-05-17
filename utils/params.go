package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var Settings Config
var BASE_DIR string

func init() {
	BASE_DIR, _ = os.Getwd() // 程序基本目录
	fmt.Println(BASE_DIR)
	configFile := filepath.Join(BASE_DIR, "server.yaml") // 配置文件
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		panic("配置文件不存在")
	}
	Settings, _ = Yaml2Struct(configFile)
}

// 日志配置
type LogSetting struct {
	Path   string `yaml:"path"`
	Name   string `yaml:"name"`
	Level  string `yaml:"level"`
	Number int8   `yaml:"number"`
}

// 配置
type Config struct {
	Port         uint16     `yaml:"port"`       // 服务运行端口
	Secert       string     `yaml:"secret"`     // jwt密钥
	Expire       uint8      `yaml:"expire"`     // 过期时间,单位:小时
	InitPassword string     `yaml:"initpasswd"` // 用户初始化或者重置密码
	IgnorePath   []string   `yaml:"ignore"`     // 忽略鉴权的路径
	Logger       LogSetting `yaml:"logger"`     // 日志设置
	FilePath     string     `yaml:"savepath"`   // 保存文件路径
}

// Yaml2Struct yaml文件转换为结构体
//
// 参数:
//
//	file string,yaml文件路径
//
// 返回:
//
//	Config: 配置结构体
//	error: 错误
func Yaml2Struct(file string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	// 默认设置
	if config.Port == 0 {
		config.Port = 80
	}
	if config.Secert == "" {
		config.Secert = "1qAz2WsX"
	}
	if config.Expire == 0 {
		config.Expire = 6
	}
	if config.InitPassword == "" {
		config.InitPassword = "zAqWe!@#123"
	}
	if config.FilePath == "" {
		config.FilePath = BASE_DIR
	}
	if config.Logger.Path == "" {
		config.Logger.Path = filepath.Join(config.FilePath, "logs")
	}
	return config, nil
}
