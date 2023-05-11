package orm

import (
	"fmt"
	"log"

	Cfg "github.com/chopper-c2-framework/c2-chopper/core/config"

	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ORMConnection struct {
	Db *gorm.DB
}

func checkMigrationError(err error) {
	if err != nil {
		log.Panicln("Error migrating to database", err)
	}

}

func CreateDB(config *Cfg.Config) (*ORMConnection, error) {
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

	return &ORMConnection{Db: db}, nil
}
