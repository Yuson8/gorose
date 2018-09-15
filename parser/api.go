package parser

import (
	"errors"
	"fizzday.com/gohouse/gorose/config"
)
// 注册解析器
var fileParsers = map[string]IParser{}

// NewFileParser 对外提供配置文件解析器接口
// fileType 文件类型
// file 文件路径
func NewFileParser(fileType, file string) (*config.DbConfig, error) {
	var ip IParser
	var err error
	if ip, err = Getter(fileType); err!=nil {
		return &config.DbConfig{}, errors.New("不支持的配置类型")
	}
	return ip.Parse(file)
}

// Getter 获取解析器
func Getter(p string) (IParser, error) {
	if pr, ok := fileParsers[p]; ok {
		return pr,nil
	}
	return nil, errors.New("解析器不存在")
}

// Register 注册解析器
func Register(p string, ip IParser) {
	fileParsers[p] = ip
	// 注册类型,方便Open()解析时区分
	config.Register(p, "file")
}