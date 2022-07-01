package migrate

import (
	// mysql_products "ppob/products/repository/mysql"
	// mysql_transaction "ppob/transaction/repository/mysql"
	// mysql_users "ppob/users/repository/mysql"
	mysql_users "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/repository/mysql/users"

	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {

	DB.AutoMigrate(&mysql_users.User{})

}
