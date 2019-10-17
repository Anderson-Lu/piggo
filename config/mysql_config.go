package config

const (
	MYSQL_SHARD_MODE_RANGE = "range"
	MYSQL_SHARD_MODE_HASH  = "hash"
)

type MysqlTableConfig struct {
	TableName   string                     `json:"table" yaml:"table"`
	ShardMode   string                     `json:"shard_mode" yaml:"shard_mode"`
	RangeConfig *MysqlShardModeRangeConfig `json:"range" yaml:"range"`
	HashConfig  *MysqlShardModeHashConfig  `json:"hash" yaml:"hash"`
}

type MysqlShardModeHashConfig struct {
	ShardSize int    `json:"shard_size" yaml:"shard_size"`
	ShardKey  string `json:"shard_key" yaml:"shard_key"`
}

type MysqlShardModeRangeConfig struct {
	ShardSize  int `json:"shard_size" yaml:"shard_size"`
	ShardRange int `json:"shard_range" yaml:"shard_range"`
}
