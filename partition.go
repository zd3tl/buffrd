package buffrd

import "context"

// Partition 定义分片结果，描述
type Partition struct {
	// id
	id int

	// loc 定义最终存在那里
	loc *location
}

type Partitioner interface {
	Pick(ctx context.Context, key string) (*Partition, error)
}
