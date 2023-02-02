package migration

import (
	"ShoppingCart/models"
	"ShoppingCart/pkg/mysql"
	"fmt"
)

func Migrate() {

	mysql.Mysql()
	
	err :=mysql.DB.AutoMigrate(
		&models.Product{},
	)

	if err != nil {
		fmt.Println("Migration Failed")
		return	
	}

	fmt.Println("Migration Success")
}