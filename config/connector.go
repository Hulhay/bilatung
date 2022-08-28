package configs

import (
	"log"
	"os"
	"sync"
	"time"

	"bilatung/internal/models"
	"bilatung/internal/repositories"
	"bilatung/internal/usecase"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var oneUc sync.Once
var uc usecase.UseCase

func GetUsecase() usecase.UseCase {
	oneUc.Do(func() {
		uc = usecase.NewUseCase(
			getRepository(),
		)
	})

	return uc
}

var repo repositories.Repositories
var oneRepo sync.Once

func getRepository() repositories.Repositories {
	oneRepo.Do(func() {
		repo = repositories.NewRepositories(getQuery())
	})

	return repo
}

var qry *gorm.DB
var oneSync sync.Once

func getQuery() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	oneSync.Do(func() {
		dbLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level; Default Silent
				Colorful:      true,        // Disable color
			},
		)

		gormConfig := &gorm.Config{
			// enhance performance config
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 dbLogger,
		}

		dsnMaster := os.Getenv("DB")
		dbMaster, errMaster := gorm.Open(mysql.Open(dsnMaster), gormConfig)
		if errMaster != nil {
			log.Panic(errMaster)
		}

		dbMaster.AutoMigrate(
			&models.Auth{},
			&models.Users{},
			&models.Quotes{},
			&models.Ratings{},
		)
		log.Print("---- Auto Migrate Success ----")

		qry = dbMaster
	})

	return qry
}

var oneInit sync.Once

func InitDB() {
	oneInit.Do(func() {
		getQuery()
	})
}
