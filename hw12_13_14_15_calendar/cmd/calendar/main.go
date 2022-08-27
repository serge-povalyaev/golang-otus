package main

import (
	"context"
	"flag"
	"github.com/serge.povalyaev/calendar/db"
	configCalendar "github.com/serge.povalyaev/calendar/internal/config/calendar"
	"github.com/serge.povalyaev/calendar/internal/repository"
	grpc2 "github.com/serge.povalyaev/calendar/internal/server/grpc"
	internalhttp "github.com/serge.povalyaev/calendar/internal/server/http"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq" // Init Database Driver
	"github.com/serge.povalyaev/calendar/internal/app"
	"github.com/serge.povalyaev/calendar/internal/logger"
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

	go func() {
		lsn, err := net.Listen("tcp", ":50051")
		if err != nil {
			logrus.Fatal(err)
		}

		grpcServer := grpc.NewServer()
		grpc2.RegisterEventServiceServer(grpcServer, grpc2.NewServer(eventRepository))
		if err := grpcServer.Serve(lsn); err != nil {
			logrus.Fatal(err)
		}
	}()

	httpServer := internalhttp.NewServer(config.Server.Host, config.Server.Port, log, eventRepository)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := httpServer.Stop(ctx); err != nil {
			log.Error("failed to stop http httpServer: " + err.Error())
		}
	}()

	log.Info("calendar is running...")

	if err := httpServer.Start(ctx); err != nil {
		log.Error("failed to start http httpServer: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
