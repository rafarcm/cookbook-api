package database

import (
	"cookbook/src/config"
	"cookbook/src/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//DBConnection -> retorna uma instancia de db
func DBConnection() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	url := config.StringConexaoBanco

	return gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})
}

// Migrate : Irá criar as tabelas do sistema no banco de dados
func Migrate(db *gorm.DB) error {
	log.Print("[Criando as tabelas do sistema]")
	return db.AutoMigrate(&model.Ingrediente{}, &model.Receita{}, &model.IngredienteReceita{})
}
