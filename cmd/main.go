package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sg4i/conf2env/pkg/converter"
)

var (
	configFile string
	outputFile string
)

func init() {
	flag.StringVar(&configFile, "conf", "config.yml", "配置文件路径")
	flag.StringVar(&outputFile, "output", ".env", "输出的环境变量文件路径")
}

func main() {
	flag.Parse()

	conv := converter.New(configFile, outputFile)
	if err := conv.Convert(); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("成功将配置从 %s 转换到 %s\n", configFile, outputFile)
}
