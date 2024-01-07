package buffrd

import "context"

// bootConfig 从配置文件启动，提供初始配置
type bootConfig struct {
	Engine string `yaml:"engine"`
}

// 内部配置中心，
type ConfigCenter interface {
	PartitionConfig(ctx context.Context) *partitionConfig
}

type config struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}

func loadConfig(ctx context.Context) error {

}

// partitionConfig 从存储engine读取配置后形成的配置，给到路由主逻辑
type partitionConfig struct {
	// domain 业务域标识，具体几层看应用设计，不冲突即可，
	// 区分业务场景，分在不同的存储engine上，例如：mysql
	// 就是db或者table，redis是不同的cluster和key
	domain string

	// partitioner 根据message得到partition
	partitioner partitioner
}
