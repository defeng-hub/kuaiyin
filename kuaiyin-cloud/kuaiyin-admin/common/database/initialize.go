package database

import (
	"os"
	"strings"
	"time"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	toolsConfig "github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	mycasbin "github.com/go-admin-team/go-admin-core/sdk/pkg/casbin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/utils"
	toolsDB "github.com/go-admin-team/go-admin-core/tools/database"
	. "github.com/go-admin-team/go-admin-core/tools/gorm/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"go-admin/common/global"
)

// Setup 配置数据库
func Setup() {
	for k := range toolsConfig.DatabasesConfig {
		setupSimpleDatabase(k, toolsConfig.DatabasesConfig[k])
	}
}

func setupSimpleDatabase(host string, c *toolsConfig.Database) {
	if global.Driver == "" {
		global.Driver = c.Driver
	}
	log.Infof("%s => %s", host, pkg.Green(c.Source))
	registers := make([]toolsDB.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := toolsDB.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.LevelForGorm()),
			},
		),
	}, opens[c.Driver])

	if err != nil {
		log.Fatal(pkg.Red(c.Driver+" connect error :"), err)
	} else {
		log.Info(pkg.Green(c.Driver + " connect success !"))
	}

	//写入同步日志文件到小时
	filePath := "logs/" + config.ApplicationConfig.Mode + "/"
	err = utils.IsNotExistMkDir(filePath)
	if err != nil {
		log.Fatal(pkg.Red(c.Driver+" 初始化同步日志文件路径失败:"), err)
		return
	}
	db.Callback().Create().After("gorm:create").Register("opt_log", func(d *gorm.DB) {
		if d.Statement.Table == "sys_job" {
			return
		}
		fileName := time.Now().Format("2006010215") + ".log"
		f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(pkg.Red(c.Driver+" 打开同步日志文件失败:"), err)
			return
		}
		defer f.Close()
		filePrefix := time.Now().Format("2006-01-02 15:04:05.000") + "\t"
		sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB { return d })
		sql = strings.ReplaceAll(sql, "\r\n", " ")
		sql = strings.ReplaceAll(sql, "\n", " ")
		_, writeErr := f.WriteString(filePrefix + sql + "\n")
		if writeErr != nil {
			log.Fatal(pkg.Red(c.Driver+" 写入同步日志文件路径失败:"), writeErr)
		}
	})
	db.Callback().Update().After("gorm:update").Register("opt_log", func(d *gorm.DB) {
		if d.Statement.Table == "sys_job" {
			return
		}
		fileName := time.Now().Format("2006010215") + ".log"
		f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(pkg.Red(c.Driver+" 打开同步日志文件失败:"), err)
			return
		}
		defer f.Close()
		filePrefix := time.Now().Format("2006-01-02 15:04:05.000") + "\t"
		sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB { return d })
		sql = strings.ReplaceAll(sql, "\r\n", " ")
		sql = strings.ReplaceAll(sql, "\n", " ")
		_, writeErr := f.WriteString(filePrefix + sql + "\n")
		if writeErr != nil {
			log.Fatal(pkg.Red(c.Driver+" 写入同步日志文件路径失败:"), writeErr)
		}
	})
	//end 写入同步日志文件

	e := mycasbin.Setup(db, "")

	sdk.Runtime.SetDb(host, db)
	sdk.Runtime.SetCasbin(host, e)
}
