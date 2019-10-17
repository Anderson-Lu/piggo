package config

type PigoQLConfig struct {
	MysqlConfig *PigoQLMysqlConfig `json:"mysql_config" yaml:"mysql_config"`
}

type PigoQLMysqlConfig struct {
	ConnectionConfigs *ConnConfig         `json:"connection" yaml:"connection"`
	TableConfigs      []*MysqlTableConfig `json:"tables" yaml:"tables"`
}
