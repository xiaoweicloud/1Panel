package migrations

import (
	"github.com/1Panel-dev/1Panel/backend/app/model"
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddSnapshotIgnore = &gormigrate.Migration{
	ID: "20240311-add-snapshot-ignore",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.Create(&model.Setting{Key: "SnapshotIgnore", Value: "*.sock"}).Error; err != nil {
			return err
		}
		return nil
	},
}

var AddDatabaseIsDelete = &gormigrate.Migration{
	ID: "20240314-add-database-is-delete",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&model.DatabaseMysql{}, &model.DatabasePostgresql{}); err != nil {
			return err
		}
		return nil
	},
}

var AddXpackHideMenu = &gormigrate.Migration{
	ID: "20240328-add-xpack-hide-menu",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.Create(&model.Setting{Key: "XpackHideMenu", Value: "{\"id\":\"1\",\"label\":\"/xpack\",\"isCheck\":true,\"title\":\"xpack.menu\",\"children\":[{\"id\":\"2\",\"title\":\"xpack.waf.name\",\"path\":\"/xpack/waf/dashboard\",\"label\":\"Dashboard\",\"isCheck\":true},{\"id\":\"3\",\"title\":\"xpack.tamper.tamper\",\"path\":\"/xpack/tamper\",\"label\":\"Tamper\",\"isCheck\":true},{\"id\":\"4\",\"title\":\"xpack.setting.setting\",\"path\":\"/xpack/setting\",\"label\":\"XSetting\",\"isCheck\":true}]}"}).Error; err != nil {
			return err
		}
		return nil
	},
}

var AddCronjobCommand = &gormigrate.Migration{
	ID: "20240403-add-cronjob-command",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&model.Cronjob{}); err != nil {
			return err
		}
		return nil
	},
}

var NewMonitorDB = &gormigrate.Migration{
	ID: "20240408-new-monitor-db",
	Migrate: func(tx *gorm.DB) error {
		var (
			bases    []model.MonitorBase
			ios      []model.MonitorIO
			networks []model.MonitorNetwork
		)
		if err := tx.Find(&bases).Error; err != nil {
			return err
		}
		if err := tx.Find(&ios).Error; err != nil {
			return err
		}
		if err := tx.Find(&networks).Error; err != nil {
			return err
		}
		if err := global.MonitorDB.AutoMigrate(&model.MonitorBase{}, &model.MonitorIO{}, &model.MonitorNetwork{}); err != nil {
			return err
		}

		_ = global.MonitorDB.Exec("DELETE FROM monitor_bases").Error
		_ = global.MonitorDB.Exec("DELETE FROM monitor_ios").Error
		_ = global.MonitorDB.Exec("DELETE FROM monitor_networks").Error

		for i := 0; i <= len(bases)/200; i++ {
			var itemData []model.MonitorBase
			if 200*(i+1) <= len(bases) {
				itemData = bases[200*i : 200*(i+1)]
			} else {
				itemData = bases[200*i:]
			}
			if err := global.MonitorDB.Create(&itemData).Error; err != nil {
				return err
			}
		}
		for i := 0; i <= len(ios)/200; i++ {
			var itemData []model.MonitorIO
			if 200*(i+1) <= len(ios) {
				itemData = ios[200*i : 200*(i+1)]
			} else {
				itemData = ios[200*i:]
			}
			if err := global.MonitorDB.Create(&itemData).Error; err != nil {
				return err
			}
		}
		for i := 0; i <= len(networks)/200; i++ {
			var itemData []model.MonitorNetwork
			if 200*(i+1) <= len(networks) {
				itemData = networks[200*i : 200*(i+1)]
			} else {
				itemData = networks[200*i:]
			}
			if err := global.MonitorDB.Create(&itemData).Error; err != nil {
				return err
			}
		}
		return nil
	},
}
