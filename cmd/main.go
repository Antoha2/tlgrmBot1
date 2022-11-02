package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Antoha2/tlgrmBot1/config"
	"github.com/Antoha2/tlgrmBot1/repository"
	"github.com/Antoha2/tlgrmBot1/service"
	trans "github.com/Antoha2/tlgrmBot1/transport"
)

func main() {

	Run()
}

func Run() {

	cfg := config.GetConfig()
	gormDB, err := initDb(cfg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	TgBotRepository := repository.NewRepository(gormDB)
	TgBotService := service.NewService(TgBotRepository)
	TgBotTransport := trans.NewWeb(TgBotService)

	TgBotTransport.StartBot()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}

func initDb(cfg *config.Config) (*gorm.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
		cfg.DB.Sslmode,
	)
	log.Println("!!!!!!!!!!!!!!!!!", connString)
	// Prep config
	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("1 failed to parse config: %v", err)
	}

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf("2 failed to create connection db: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbx,
	}), &gorm.Config{})

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf("3 error to ping connection pool: %v", err)
	}
	log.Printf("Подключение к базе данных на http://127.0.0.1:%d\n", cfg.DB.Port)
	return gormDB, nil
}
