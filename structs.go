package buffrd

type message struct {
	// Domain 目标分片集合的标识，利用http传输，非特殊的
	// client-server协议，需要在消息中指定。
	Domain string `json:"domain"`

	// Key 用于分片策略
	Key string `json:"key"`

	// Data 需要业务处理的任务信息
	Data string `json:"data"`

	// Timestamp 预期处理的时间，按照消息的前后顺序处理
	// 应用设计场景设定：
	// 实时任务
	// 消息不接受长时间延迟，包括评估后和不评估两种case，
	// 不评估这个case比较有意思，产品起步初期，或者可以不考虑成本，
	// 都直接使用扩容手段让任务不堆积处理，同时也缺少对消息处理质
	// 量的的统计。
	// 延时任务
	// 适用于可以预估周期任务的场景，且任务可以接受较长延时，应用方
	// 需要规划不同业务场景消息处理的时间边界，并能够明确感知消息完结
	// ，防止不同业务域消息重叠在同样的任务处理集群上，造成资源竞争。
	// TODO 优先实时的场景
	Timestamp int64 `json:"timestamp"`
}

type schemaType string

const (
	mysqlSchema schemaType = "mysql"
)

type location struct {
	// schema 区分不同存储引擎
	schema schemaType
}

type mysqlLocation struct {
	DBName    string `json:"dbName"`
	TableName string `json:"tableName"`
}
