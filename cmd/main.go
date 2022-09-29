package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bin "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"net/http"
	"secret-service/app"
	"secret-service/config"
	"secret-service/infrastructure/handlers"
	"secret-service/infrastructure/repository"
	"secret-service/migrations"
)

func main() {
	var config config.Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
		return
	}

	migrateDB(&config)

	db := connectDB(&config)
	repoUser := repository.NewUserRepositoryDb(db)

	serviceUser := app.NewUserService(repoUser)
	secretService := app.NewSecretService(repoUser)

	handlerUser := handlers.NewUserHandler(serviceUser)
	handlerSecret := handlers.NewSecretHandler(secretService)

	go app.CreatedTimeCompare(db)

	e := echo.New()
	e.POST("/add_user", handlerUser.CreateUser)
	e.POST("/secret", handlerSecret.CreateSecret)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ListenPort), e))
}

func connectDB(conf *config.Config) (db *sql.DB) {
	psqlConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.ConfigDataBase.Host, conf.ConfigDataBase.Port, conf.ConfigDataBase.User, conf.ConfigDataBase.Password, conf.ConfigDataBase.NameDataBase)
	db, err := sql.Open("postgres", psqlConnStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func migrateDB(conf *config.Config) {
	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require",
		conf.ConfigDataBase.User,
		conf.ConfigDataBase.Password,
		conf.ConfigDataBase.Host,
		conf.ConfigDataBase.Port,
		conf.ConfigDataBase.NameDataBase,
	)

	source := bin.Resource(migrations.AssetNames(), migrations.Asset)
	driver, err := bin.WithInstance(source)
	if err != nil {
		log.Fatal(err)
	}
	migration, err := migrate.NewWithSourceInstance("go-bindata", driver, databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err = migration.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Warn(err)
		} else {
			log.Fatal(err)
		}
	}
}
