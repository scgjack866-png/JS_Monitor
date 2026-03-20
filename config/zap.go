package config

// zap配置
type Zap struct {
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}
