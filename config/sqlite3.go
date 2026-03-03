package config

// Sqlite3 配置参数
type Sqlite3 struct {
	Path string
}

// Sqlite3 数据库连接串
func (a Sqlite3) DSN() string {
	return a.Path
}
