package buffrd

import "context"

// Partition 定义分片结果，描述
type partition struct {
	// id
	id int

	// loc 定义最终存在那里
	loc *location
}

type partitioner interface {
	partition(ctx context.Context, key string) (*partition, error)
}
type mysqlHashPartitioner struct {
	db    string
	table string
}

// partition
// key 消息的唯一标识，用于做hash分布的源数据，具体业务域的分片集合需要在client初始化时提供
func (p *mysqlHashPartitioner) partition(ctx context.Context, domain string, key string) (*partition, error) {
	// TODO 利用domain找到分片集合，在集合内做映射即可
}
