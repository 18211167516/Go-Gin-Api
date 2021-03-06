package initialize

//https://gorm.io/docs/gorm_config.html
import (
	"go-api/global"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func GormMysql() *gorm.DB {
	m := global.CF.Mysql
	dsn := m.MysqlUser + ":" + m.MysqlPassword + "@tcp(" + m.MysqlHost + ")/" + m.MysqlName + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), GromConfig()); err != nil {
		global.LOG.Error("Mysql启动异常", dsn, err)
		os.Exit(0)
		return nil
	} else {
		db.Use(dbresolver.Register(dbresolver.Config{
			Sources: []gorm.Dialector{
				mysql.Open("root:123456@tcp(192.168.99.100:3306)/test"),
			},
			Replicas: []gorm.Dialector{
				mysql.Open("root:123456@tcp(192.168.99.100:3306)/test2"),
				mysql.Open("root:123456@tcp(192.168.99.100:3306)/test3"),
			},
			// sources/replicas 负载均衡策略
			Policy: dbresolver.RandomPolicy{},
		}).SetConnMaxIdleTime(10).
			SetConnMaxLifetime(m.MaxLifetime).
			SetMaxOpenConns(10).
			SetMaxIdleConns(2),
		)

		db.Clauses(dbresolver.Read)
		//sqlDB, _ := db.DB()
		//sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
		//sqlDB.SetMaxOpenConns(100)              //最大数据库链接数
		//sqlDB.SetConnMaxLifetime(m.MaxLifetime) //数据库链接最大生存时间             //true  打印Log
		return db
	}

}

func GromConfig() *gorm.Config {
	//m := global.CF.Mysql
	c := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   getGromLogger(),
		/* NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.MysqlPrefix, // 表前缀
			SingularTable: true,          // 使用单数表名
		}, */
	}
	return c
}

func getGromLogger() logger.Interface {
	LogLevel := logger.Silent
	if global.CF.Mysql.LogMode {
		LogLevel = logger.Info
	}

	//log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	newLogger := logger.New(
		global.LOG,
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      LogLevel,    // Log level
			Colorful:      false,       // Disable color
		},
	)

	return newLogger
}
