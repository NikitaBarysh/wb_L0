package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/NikitaBarysh/wb_L0/cmd/cfg"
	cache2 "github.com/NikitaBarysh/wb_L0/internal/cache"
	"github.com/NikitaBarysh/wb_L0/internal/handler"
	"github.com/NikitaBarysh/wb_L0/internal/repository"
	"github.com/NikitaBarysh/wb_L0/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := cfg.NewConfig()
	logrus.Info("project config: ", cfg)

	cache := cache2.NewCache()
	db, err := repository.NewPostgresDB(cfg.DataBaseDSN)
	logrus.Info("database path: ", cfg.DataBaseDSN)
	if err != nil {
		logrus.Error("main: NewPostgresDB: %w", err)
	}

	storage := repository.NewRepository(db)
	storage.OrderAutoMigrate()
	newService := service.NewService(cache, cfg, storage)
	newService.Order.RestoreCache()
	handlers := handler.NewHandler(newService)
	router := handler.NewRouter(handlers)
	chiRouter := chi.NewRouter()
	chiRouter.Mount("/", router.Register())

	go newService.NATSProducer.Publish()
	newService.Nats.SubscribeToNATS()

	go func() {
		err = http.ListenAndServe(cfg.ServerEndpoint, chiRouter)
	}()
	logrus.Info("server started with port: ", cfg.ServerEndpoint)

	termSig := make(chan os.Signal, 1)
	signal.Notify(termSig, syscall.SIGTERM, syscall.SIGINT)
	sig := <-termSig

	logrus.Info("Shutting Down ", sig.String())

}
