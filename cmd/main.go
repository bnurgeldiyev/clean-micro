package main

import (
	"clean-micro/internal/adapters/api"
	"clean-micro/internal/adapters/service/cache"
	"clean-micro/internal/composites"
	"clean-micro/internal/config"
	"clean-micro/pkg/logging"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	err := config.ReadConfig("../pkg/config.json")
	if err != nil {
		log.WithError(err).Panic("error reading config file")
	}

	logging.SetupLogger()

	setupServer(signalChan, config.Conf)
}

func setupServer(signalChan chan os.Signal, conf *config.Config) {

	listener, err := net.Listen("tcp", conf.ListenAddress)
	if err != nil {
		log.WithError(err).Panic("error in net.Listen")
	}

	postgresComposite, err := composites.NewPostgresComposite(conf.DbConn, conf.DbMaxConn)
	if err != nil {
		log.WithError(err).Panic("error in composites.NewPostgresComposite()")
	}

	redisService := cache.NewRedisService(conf.RedisConn, conf.RedisDB, 8, 256)

	userComposite, err := composites.NewUserComposite(postgresComposite, redisService)
	if err != nil {
		log.WithError(err).Panic("composites.NewUserComposite()")
	}

	grpcServer := grpc.NewServer()
	api.RegisterAuthServer(grpcServer, userComposite.Handler)

	go startServer(grpcServer, listener)

	fmt.Println(conf.DbConn)
	fmt.Println("<--START-SERVER-->", conf.ListenAddress)
	for {
		select {
		case sig := <-signalChan:
			switch sig {
			case os.Interrupt:
				fmt.Println("interrupt")
				return
			case os.Kill:
				fmt.Println("kill")
				return
			}
		}
	}
}

func startServer(grpcServer *grpc.Server, listener net.Listener) {
	err := grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
