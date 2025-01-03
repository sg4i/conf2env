# Conf2Env

一个简单的工具，用于将配置文件（YAML）转换为环境变量文件（.env）。

## 功能特点

- 支持多种配置文件格式（YAML、JSON等）
- 自动将嵌套配置转换为扁平化的环境变量
- 支持自定义输入和输出文件路径

## 安装

```bash
go install github.com/sg4i/conf2env/cmd@latest
```

或者从源码构建：

```bash
make build
```

## 使用方法

基本用法：

```bash
conf2env -conf config.yml -output .env
```

参数说明：

- `-conf`: 配置文件路径（默认：config.yml）
- `-output`: 输出的环境变量文件路径（默认：.env）

## 示例

输入文件 `config.yml`：

```yaml
app:
  name: myapp
  port: 8080
database:
  host: localhost
  port: 5432
```

生成的 `.env` 文件：

```env
APP_NAME=myapp
APP_PORT=8080
DATABASE_HOST=localhost
DATABASE_PORT=5432
```

## 许可证

MIT License
