package config

// Zap zap配置
// Level: 自定义日志级别，支持 debug/info/warn/error 等 zap 标准级别。
// Debug: 开启后会额外输出到 stdout，并联动开启 Gin/Gorm 的调试行为。
type Zap struct {
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Level      string
	Debug      bool
}
