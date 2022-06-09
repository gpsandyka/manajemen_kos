package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection() (*gorm.DB, error) {
	// dsn := "host=localhost user=makkoswebapp password=1w3r5y7i dbname=makkosdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=ec2-34-198-186-145.compute-1.amazonaws.com " +
		"user=ttcvhxmmpcrxtb " +
		"password=e5bc24157fd742c3ee4476115e8b6964a03c40706780966b53d1c340f4ecbd18 " +
		"dbname=dev6jnegjtkqj3 " +
		"port=5432 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
