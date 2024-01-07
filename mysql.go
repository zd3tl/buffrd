package buffrd

import (
	"context"
	"encoding/json"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// MySQLConfigKV mysql配置表单条记录
type mysqlConfigKV struct {
	Id int `db:"id"`

	// Domain 业务域
	Domain string `db:"domain"`

	// Key 需要做到区分domain
	Key   string `db:"key"`
	Value string `db:"value"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	// Version 标记版本
	Version string `db:"version"`
	// Enable 启用，多版本场景或者下架配置场景都可以使用
	Enable bool `db:"enable"`
}

type strategyNameType string

const (
	partitionStrategyMySQLSimple     strategyNameType = "PartitionStrategyMySQLSimple"
	partitionStrategyMySQLMultiTable strategyNameType = "PartitionStrategyMySQLMultiTable"
	partitionStrategyMySQLMultiDB    strategyNameType = "PartitionStrategyMySQLMultiDB"
)

type mysqlPartitionConfig struct {
	Strategy *mysqlHashPartitionStrategy `json:"strategy"`
}

// mysqlHashPartitionConfig mysql场景策略：单库单表多p、单库多表多p，多库多表多p
type mysqlHashPartitionStrategy struct {
	// Name strategy名称
	Name strategyNameType `json:"name"`

	// PartitionCount 分片数量
	// hash场景只需要知道范围：(0, PartitionCount]，message的Key映射到这个范围即可
	PartitionCount int `json:"partitionCount"`
}

// mysqlConfigCenter
type mysqlConfigCenter struct {
	db *sqlx.DB
}

func newMySQLConfigCenter(cfg *bootConfig) (configCenter, error) {
	db, err := sqlx.Connect("mysql", cfg.DSN)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 保持internal不开放，设计上不会有巨量访问
	// TODO 可以开放 + default的模式，让这块优雅起来
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	return &mysqlConfigCenter{db: db}, nil
}

func (cc *mysqlConfigCenter) getPartitioner(_ context.Context, domain string) (Partitioner, error) {
	var kv mysqlConfigKV
	if err := cc.db.QueryRowx("SELECT * FROM config WHERE domain=? AND key='Pick' AND enable=true", domain).Scan(&kv); err != nil {
		return nil, errors.WithStack(err)
	}

	// 加载特定domain的partition配置
	var cfg mysqlPartitionConfig
	if err := json.Unmarshal([]byte(kv.Value), &cfg); err != nil {
		return nil, errors.WithStack(err)
	}
	return &mysqlHashPartitioner{cfg: &cfg}, nil
}

type mysqlHashPartitioner struct {
	cfg *mysqlPartitionConfig
}

// Pick
// key 消息的唯一标识，用于做hash分布的源数据，具体业务域的分片集合需要在client初始化时提供
func (p *mysqlHashPartitioner) Pick(ctx context.Context, key string) (*Partition, error) {
	// TODO 利用domain找到分片集合，在集合内做映射即可
	return nil, nil
}
