package migration

import (
	"fmt"
	"golang-restfull-api/db"
	"log"
)

func MigrationHandler() {
	err := db.DBInit().AutoMigrate(&Biodata{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database migrated")
}