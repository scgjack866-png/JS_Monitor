package config

import (
	"github.com/spf13/viper"
)

// Config 配置参数
type Config struct {
	Grafana    Grafana
	Prometheus Prometheus
	Web        Web
	Gorm       Gorm
	MySQL      MySQL
	Sqlite3    Sqlite3
	Zap        Zap
}

func LoadConfig(fpath string) (c *Config, err error) {
	v := viper.New()
	v.SetConfigFile(fpath)
	v.SetConfigType("yaml")
	v.SetDefault("gorm.enable_sql_log", true)

	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}

	c = &Config{}
	c.Grafana.ApiUrl = v.GetString("grafana.api_url")
	c.Grafana.Authorization = v.GetString("grafana.authorization")
	c.Grafana.PrometheusUid = v.GetString("grafana.prometheusUid")
	c.Grafana.ConfigPath = v.GetString("grafana.configPath")
	c.Grafana.SnapshotDomain = v.GetString("grafana.snapshotDomain")
	c.Prometheus.ApiUrl = v.GetString("prometheus.api_url")
	c.Web.StaticPath = v.GetString("web.static_path")
	c.Web.Domain = v.GetString("web.domain")
	c.Web.Port = v.GetInt("web.port")
	c.Web.ReadTimeout = v.GetInt("web.read_timeout")
	c.Web.WriteTimeout = v.GetInt("web.write_timeout")
	c.Web.IdleTimeout = v.GetInt("web.idle_timeout")
	c.MySQL.Host = v.GetString("mysql.host")
	c.MySQL.Port = v.GetInt("mysql.port")
	c.MySQL.User = v.GetString("mysql.user")
	c.MySQL.Password = v.GetString("mysql.password")
	c.MySQL.DBName = v.GetString("mysql.db_name")
	c.MySQL.Parameters = v.GetString("mysql.parameters")
	c.Sqlite3.Path = v.GetString("sqlite3.path")
	c.Gorm.Debug = v.GetBool("gorm.debug")
	c.Gorm.EnableSQLLog = v.GetBool("gorm.enable_sql_log")
	c.Gorm.DBType = v.GetString("gorm.db_type")
	c.Gorm.MaxLifetime = v.GetInt("gorm.max_lifetime")
	c.Gorm.MaxOpenConns = v.GetInt("gorm.max_open_conns")
	c.Gorm.MaxIdleConns = v.GetInt("gorm.max_idle_conns")
	c.Gorm.TablePrefix = v.GetString("gorm.table_prefix")
	// zap 日志相关配置：文件滚动、日志级别、debug 开关。
	c.Zap.FileName = v.GetString("zap.fileName")
	c.Zap.MaxSize = v.GetInt("zap.maxSize")
	c.Zap.MaxAge = v.GetInt("zap.maxAge")
	c.Zap.MaxBackups = v.GetInt("zap.maxBackup")
	c.Zap.Level = v.GetString("zap.level")
	c.Zap.Debug = v.GetBool("zap.debug")
	return
}
