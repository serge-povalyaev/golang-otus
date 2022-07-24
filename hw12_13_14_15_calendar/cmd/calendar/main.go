package main

import (
	"context"
	"flag"
	"github.com/serge.povalyaev/calendar/db"
	configCalendar "github.com/serge.povalyaev/calendar/internal/config/calendar"
	"github.com/serge.povalyaev/calendar/internal/repository"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/serge.povalyaev/calendar/internal/app"
	"github.com/serge.povalyaev/calendar/internal/logger"
	internalhttp "github.com/serge.povalyaev/calendar/internal/server/http"

	_ "github.com/lib/pq" // Init Database Driver
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config/calendar.yaml", "Путь до файла конфигурации")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := configCalendar.ReadConfig(configPath)

	// Настройка логгера
	log := logger.New(config.Logger.Level, config.Logger.FilePath)

	// Настройка подключения к БД
	dsn := db.CreateDSN(
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Pass,
		config.DB.Name,
	)
	connection := db.Connect(dsn)

	// Настройка репозитория
	eventRepository := repository.GetEventRepository(config.RepositoryType, connection)

	_ = app.New(*log, eventRepository, connection)

	server := internalhttp.NewServer(config.Server.Host, config.Server.Port, log, eventRepository)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			log.Error("failed to stop http server: " + err.Error())
		}
	}()

	log.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		log.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
