package orm

import (
	"fmt"
	"log"

	Cfg "github.com/chopper-c2-framework/c2-chopper/core/config"

	"github.com/chopper-c2-framework/c2-chopper/server/domain/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IORMConnection interface {
	CreateDB(config *Cfg.Config) error
}

type ORMConnection struct {
	db *gorm.DB
}

func checkMigrationError(err error) {
	if err != nil {
		log.Panicln("Error migrating to database", err)
	}

}

func (conn *ORMConnection) CreateDB(config *Cfg.Config) error {
	db, err := gorm.Open(sqlite.Open(config.ServerDb), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("[+] Created DB:", config.ServerDb)

	// Migrate the schema
	checkMigrationError(db.AutoMigrate(&entity.UserModel{}))
	checkMigrationError(db.AutoMigrate(&entity.TeamModel{}))
	checkMigrationError(db.AutoMigrate(&entity.TaskModel{}))
	checkMigrationError(db.AutoMigrate(&entity.TaskResultModel{}))
	checkMigrationError(db.AutoMigrate(&entity.ListenerModel{}))

	fmt.Println("[+] Migrated Models.")

	conn.db = db
	return nil
}
