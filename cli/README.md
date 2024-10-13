# Blog Helper
gopiblog的管理命令行管理工具

目前有统计信息, 上传管理, 回复查询的功能

## Example
```bash
go build
./piblog.exe -h
```

在程序没有检测到二进制文件的同一目录下的配置文件/博文模板文件时,会自动创建

## Configure
Connect to grpc server

由于go-zero的`zrpc/client.go`中`NewClient()`在build rpc target时可能引入grpc尚未支持的uri syntax `etcd`,所以建议只配置grpc server address