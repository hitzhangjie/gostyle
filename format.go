package gostyle

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const (
	GO     = "go"
	JAVA   = "java"
	CPP    = "cpp"
	PYTHON = "python"
)

type Formatter func(source string) error

var (
	formatters = map[string]Formatter{
		GO: GoFormat,
	}
)

// FormatSource 原地格式化代码
func FormatSource(source, lang string) error {

	lang = strings.ToLower(lang)
	formatter, ok := formatters[lang]
	if !ok {
		return errors.New("invalid lang")
	}
	return formatter(source)
}

// FormatSourceDir 原地格式化go代码目录
func FormatSourceDir(dir, lang string) error {

	lang = strings.ToLower(lang)
	formatter, ok := formatters[lang]
	if !ok {
		return errors.New("invalid lang")
	}

	err := filepath.Walk(dir, func(fpath string, info os.FileInfo, err error) error {
		if strings.HasSuffix(fpath, ".go") && !info.IsDir() {
			return formatter(fpath)
		}
		return nil
	})
	return err
}

// Register 为语言lang注册对应的格式化方法
func Register(lang string, formatter Formatter) {
	formatters[lang] = formatter
}
