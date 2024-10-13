package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	HostRpc zrpc.RpcClientConf // access to rpc 'host.rpc'
}

var GlobalConfig Config
var exepath string

const configFileName string = "conf.yaml"

func init() {
	exepath, _ = os.Executable()
	GuaranteeConfigure()
	conf.MustLoad(
		fmt.Sprintf("%s/%s", filepath.Dir(exepath), "conf.yaml"), &GlobalConfig)
}

const confFile = `HostRpc:
  # Etcd:
  #   Hosts:
  #     - http://localhost:2379
  #   Key: host.rpc0
  Target: "127.0.0.1:8081"
  App: "MyCli"
  Token: "SoEasyToken"
`

func GuaranteeConfigure() {

	confPath := fmt.Sprintf("%s/%s", filepath.Dir(exepath), configFileName)
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		fmt.Println("第一次创建配置文件")
		obj, err := os.Create(confPath)
		if err != nil {
			panic(err)
		}
		_, err = obj.WriteString(confFile)
		if err != nil {
			panic(err)
		}
	}
}
