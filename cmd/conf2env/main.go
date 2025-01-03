package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sg4i/conf2env/pkg/converter"
)

var (
	Version = "dev" // 由编译时注入

	configFile string
	outputFile string
)

func init() {
	flag.StringVar(&configFile, "conf", "config.yml", "配置文件路径")
	flag.StringVar(&outputFile, "output", ".env", "输出的环境变量文件路径")
	flag.Parse()
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Printf("conf2env version %s\n", Version)
		return
	}

	conv := converter.New(configFile, outputFile)
	if err := conv.Convert(); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("成功将配置从 %s 转换到 %s\n", configFile, outputFile)
}
