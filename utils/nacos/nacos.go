package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NewNacosConfig(ipAddr, namespaceId, dataId string) []byte {
	sc := []constant.ServerConfig{{
		IpAddr:      ipAddr,
		Port:        8848,
		ContextPath: "/nacos",
		Scheme:      "http",
	}}
	cc := constant.ClientConfig{
		NamespaceId:         namespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		panic(err)
	}
	return []byte(content)
}
