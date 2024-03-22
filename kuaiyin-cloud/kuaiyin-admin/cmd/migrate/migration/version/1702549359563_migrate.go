package version

import (
	"runtime"

	"gorm.io/gorm"

	"go-admin/cmd/migrate/migration"
	"go-admin/common/models"
	common "go-admin/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1702549359563Test)
}

func _1702549359563Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		// TODO: 例如 新增表结构 使用过程中请删除此段代码
		err := tx.Debug().Migrator().AutoMigrate(
			new(TbSyncLog1702549359563),
		)
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}

type TbSyncLog1702549359563 struct {
	models.Model

	Env     string `json:"env" gorm:"type:varchar(128);comment:环境"`
	LogName string `json:"logName" gorm:"type:varchar(128);comment:文件名"`
	Status  int64  `json:"status" gorm:"type:tinyint(1);comment:同步状态"`
	models.ModelTime
	models.ControlBy
}

func (TbSyncLog1702549359563) TableName() string {
	return "tb_sync_log"
}
