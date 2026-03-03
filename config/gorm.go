package config

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool
	DBType       string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
	DSN          string
}
