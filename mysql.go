package buffrd

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// MySQLConfigKV mysql配置表单条记录
type mysqlConfigKV struct {
	Id int `db:"id"`

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

// mysqlConfigCenter
type mysqlConfigCenter struct {
	db *sqlx.DB
}

func newMySQLConfigCenter() (ConfigCenter, error) {
	db, err := sqlx.Connect("mysql", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return &mysqlConfigCenter{
		db: db,
	}, nil
}

func (cc *mysqlConfigCenter) PartitionConfig(ctx context.Context) *partitionConfig {
	return nil
}
