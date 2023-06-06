package main

import (
	"driver/config"
	"driver/migration"
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println("mragration")
	doMigration()
}

var myConfig = config.GetConfig()

func doMigration() {
	migration := migration.NewMigration()
	target := myConfig.MigrationTarget
	if target == "latest" {
		migration.Up()
	} else {
		targetNum, err := strconv.ParseUint(myConfig.MigrationTarget, 10, 0)
		if err != nil {
			log.Fatal(err)
		}
		migration.To(uint(targetNum))
	}
}
