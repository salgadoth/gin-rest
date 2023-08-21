package database

import (
	"log"

	"github.com/salgadoth/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComDB() {
	connString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connString))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados.")
	}

	DB.AutoMigrate(&models.Aluno{})
}
