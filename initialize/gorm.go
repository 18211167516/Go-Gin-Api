package initialize

//https://gorm.io/docs/gorm_config.html
import (
	"fmt"
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
	dsn := getMasterDsn()
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	global.LOG.Infoln("slave", global.VP.GetString("mysql.slave.0.DBName"))
	if db, err := gorm.Open(mysql.New(mysqlConfig), GromConfig()); err != nil {
		global.LOG.Error("Mysql启动异常", dsn, err)
		os.Exit(0)
		return nil
	} else {
		if len := global.VP.GetInt("mysql.global.slave"); len > 0 {
			Replicas := make([]gorm.Dialector, len)
			for i := 0; i < len; i++ {
				Replicas = append(Replicas, mysql.Open(getSlaveDsn(i)))
			}
			db.Use(dbresolver.Register(dbresolver.Config{
				Sources: []gorm.Dialector{
					mysql.Open(dsn),
				},
				Replicas: Replicas,
				// sources/replicas 负载均衡策略
				Policy: dbresolver.RandomPolicy{},
			}).SetConnMaxIdleTime(global.VP.GetDuration("mysql.global.MaxIdleTime")).
				SetConnMaxLifetime(global.VP.GetDuration("mysql.global.MaxLifetime")).
				SetMaxOpenConns(global.VP.GetInt("mysql.global.MaxOpenConns")).
				SetMaxIdleConns(global.VP.GetInt("mysql.global.MaxIdleConns")),
			)
		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetConnMaxIdleTime(global.VP.GetDuration("mysql.global.MaxIdleTime")) //数据库连接最大空闲时间
			sqlDB.SetConnMaxLifetime(global.VP.GetDuration("mysql.global.MaxLifetime")) //数据库连接可复用的最大时间
			sqlDB.SetMaxOpenConns(global.VP.GetInt("mysql.global.MaxOpenConns"))        //最大数据库链接数
			sqlDB.SetMaxIdleConns(global.VP.GetInt("mysql.global.MaxIdleConns"))        //最大空闲连接数
		}

		//db.Clauses(dbresolver.Read) 使用读链接 从库
		//db.Clauses(dbresolver.Write) 使用写链接 主库

		return db
	}

}

//获取主库dsn
func getMasterDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		global.VP.GetString("mysql.master.User"),
		global.VP.GetString("mysql.master.Password"),
		global.VP.GetString("mysql.master.Host"),
		global.VP.GetString("mysql.master.DBName"),
		global.VP.GetString("mysql.master.Config"),
	)
}

//获取从库DSN
func getSlaveDsn(i int) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		global.VP.GetString(fmt.Sprintf("mysql.Slave.%d.User", i)),
		global.VP.GetString(fmt.Sprintf("mysql.Slave.%d.Password", i)),
		global.VP.GetString(fmt.Sprintf("mysql.Slave.%d.Host", i)),
		global.VP.GetString(fmt.Sprintf("mysql.Slave.%d.DBName", i)),
		global.VP.GetString(fmt.Sprintf("mysql.Slave.%d.Config", i)),
	)
}

func GromConfig() *gorm.Config {
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
	if global.VP.GetBool("mysql.global.LogMode") {
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
