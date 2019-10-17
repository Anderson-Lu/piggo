package config

type ConnConfig struct {
	Host    string `json:"host" yaml:"host"`
	Port    string `json:"port" yaml:"port"`
	User    string `json:"user" yaml:"user"`
	Pass    string `json:"pass" yaml:"pass"`
	MaxConn int    `json:"max_conn" yaml:"max_conn"`
	MaxIdle int    `json:"max_idle" yaml:"max_idle"`
}
