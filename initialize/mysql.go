package initialize

import (
	"OperationAndMonitoring/config"
	"OperationAndMonitoring/mysql/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(config *config.Config) {
	var gdb *gorm.DB
	var err error

	if config.Gorm.DBType == "mysql" {
		config.Gorm.DSN = config.MySQL.DSN()
	} else if config.Gorm.DBType == "sqlite3" {
		config.Gorm.DSN = config.Sqlite3.DSN()
	}
	mysqlConfig := mysql.Config{
		DSN:                       config.Gorm.DSN, // DSN data source name
		DefaultStringSize:         191,             // string 类型字段的默认长度
		SkipInitializeWithVersion: false,           // 根据版本自动配置
	}

	gormConfig := &gorm.Config{}

	// debug 模式下提升为 Info，便于查看 SQL 执行过程；默认仅记录错误。
	logLevel := logger.Error
	if config.Gorm.Debug || config.Zap.Debug {
		logLevel = logger.Info
	}
	gormConfig.Logger = db.NewGormLogger(logLevel, config.Gorm.EnableSQLLog)

	//newlogger := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.Lmsgprefix)), logger.Config{
	//	SlowThreshold: 200 * time.Millisecond,
	//	LogLevel:      logger.Info,
	//	Colorful:      true,
	//})

	gdb, err = gorm.Open(mysql.New(mysqlConfig), gormConfig)

	if err != nil {
		panic(err)
	}
	db.DB = gdb
}

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}
