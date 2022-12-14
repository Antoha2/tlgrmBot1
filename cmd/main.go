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
	"github.com/Antoha2/tlgrmBot1/internal/geokoder"
	"github.com/Antoha2/tlgrmBot1/internal/meteo/providers/gismeteo"
	yandeX "github.com/Antoha2/tlgrmBot1/internal/meteo/providers/yandeX"
	http "github.com/Antoha2/tlgrmBot1/pkg/http/transport"
	repository "github.com/Antoha2/tlgrmBot1/repository"
	service "github.com/Antoha2/tlgrmBot1/service/windService"
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

	TgBotGeokoder := geokoder.NewGeokoder()
	TgBotMeteoYandex := yandeX.NewYandex(cfg)
	TgBotMeteoGismeteo := gismeteo.NewGismeteo()

	TgBotService := service.NewService(TgBotRepository, TgBotMeteoYandex, TgBotMeteoGismeteo, TgBotGeokoder)
	TgBotTransport := trans.NewWeb(TgBotService, cfg)
	HTTPTransport := http.NewHTTP(TgBotService, cfg)

	go TgBotTransport.StartBot()
	go HTTPTransport.StartHTTP()

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
	if err != nil {
		return nil, fmt.Errorf("3 gorm.Open(): %v", err)
	}

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf("4 error to ping connection pool: %v", err)
	}
	log.Printf("?????????????????????? ?? ???????? ???????????? ???? http://127.0.0.1:%d\n", cfg.DB.Port)
	return gormDB, nil
}
