# 项目依赖

操作系统：Ubuntu24.04

GoLang版本： go1.26.2

# 依赖安装

1. 下载 Go 1.26.2 安装包

   ```shell
   wget https://dl.google.com/go/go1.26.2.linux-amd64.tar.gz
   ```

2. 解压安装包到系统标准目录

   ```shell
   sudo tar -C /usr/local -xzf go1.26.2.linux-amd64.tar.gz
   ```

3. 配置环境变量

   ```shell
   sudo nano /etc/profile
   ```

   在文件末尾添加：

   ```shell
   export PATH=$PATH:/usr/local/go/bin
   ```

   执行：

   ```sh
   source /etc/profile
   ```

4. 验证

   ```shell
   go version
   ```

   显示

   ```shell
   go version go1.26.2 linux/amd64
   ```

# 启动项目

解压缩并进入项目目录：

```shell
tar -xzvf simple-block-api.tar.gz
cd simple-block-api/
```

安装依赖包

```shell
go mod tidy
```

启动web服务

```shell
go run main.go
```

