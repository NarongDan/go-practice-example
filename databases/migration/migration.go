package main

import (
	"fmt"
	"tutorial/config"
	"tutorial/databases"
	"tutorial/entities"

	"gorm.io/gorm"
)

func main() {

	conf := config.ConfigGetting()

	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.Connect())

	tx := db.Connect().Begin()

	// playerMigration(db)
	// adminMigration(db)
	// itemMigration(db)
	// playerCoinMigration(db)
	// inventoryMigration(db)
	// purchaseHistoryMigration(db)

	playerMigration(tx)
	adminMigration(tx)
	itemMigration(tx)
	playerCoinMigration(tx)
	inventoryMigration(tx)
	purchaseHistoryMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

// func playerMigration(db databases.Database) {
// 	db.Connect().Migrator().CreateTable(&entities.Player{})
// }

// func adminMigration(db databases.Database) {
// 	db.Connect().Migrator().CreateTable(&entities.Admin{})
// }

// func itemMigration(db databases.Database) {
// 	db.Connect().Migrator().CreateTable(&entities.Item{})
// }

// func playerCoinMigration(db databases.Database) {
// 	db.Connect().Migrator().CreateTable(&entities.PlayerCoin{})
// }

// func inventoryMigration(db databases.Database) {
// 	db.Connect().Migrator().CreateTable(&entities.Inventory{})
// }

// func purchaseHistoryMigration(db databases.Database) {
// 	db.Connect().Migrator().CreateTable(&entities.PurchaseHistory{})
// }

func playerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Player{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
