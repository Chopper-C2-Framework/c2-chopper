package orm

import (
	"fmt"

	Cfg "github.com/chopper-c2-framework/c2-chopper/core/config"

	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ORMConnection struct {
	Db *gorm.DB
}

func CreateDB(config *Cfg.Config) (*ORMConnection, error) {
	db, err := gorm.Open(sqlite.Open(config.ServerDb), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("[+] Created DB:", config.ServerDb)

	// Migrate the schema
	db.AutoMigrate(&entity.UserModel{})
	db.AutoMigrate(&entity.TeamModel{})
	db.AutoMigrate(&entity.TaskModel{})
	db.AutoMigrate(&entity.TaskResultModel{})
	db.AutoMigrate(&entity.ListenerModel{})

	fmt.Println("[+] Migrated Models.")

	conn := ORMConnection{Db: db}
	return &conn, nil
}
