package config

// 站点配置参数
type Web struct {
	Domain       string
	StaticPath   string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}
