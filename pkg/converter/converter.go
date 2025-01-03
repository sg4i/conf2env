package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// Converter 配置转换器
type Converter struct {
	configFile string
	outputFile string
}

// New 创建新的转换器实例
func New(configFile, outputFile string) *Converter {
	return &Converter{
		configFile: configFile,
		outputFile: outputFile,
	}
}

// Convert 执行配置转换
func (c *Converter) Convert() error {
	// 设置配置文件
	ext := filepath.Ext(c.configFile)
	if ext == "" {
		ext = ".yml"
	}
	viper.SetConfigType(strings.TrimPrefix(ext, "."))
	viper.SetConfigFile(c.configFile)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 创建输出文件
	envFile, err := os.Create(c.outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %w", err)
	}
	defer envFile.Close()

	// 获取所有配置并写入
	return c.writeConfig(envFile, viper.AllSettings(), "")
}

// writeConfig 递归写入配置
func (c *Converter) writeConfig(file *os.File, config map[string]interface{}, prefix string) error {
	for key, value := range config {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		// 将键转换为大写并替换点为下划线
		envKey := strings.ToUpper(strings.ReplaceAll(fullKey, ".", "_"))

		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			if err := c.writeConfig(file, value.(map[string]interface{}), fullKey); err != nil {
				return err
			}
		default:
			if _, err := fmt.Fprintf(file, "%s=%v\n", envKey, value); err != nil {
				return fmt.Errorf("写入配置失败: %w", err)
			}
		}
	}
	return nil
}
