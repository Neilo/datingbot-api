package main

import (
	"database/sql"
	"github.com/pressly/goose"
	"os"
	"os/signal"
	"syscall"

	"github.com/brotherhood228/dating-bot-api/internal/health"
	_ "github.com/brotherhood228/dating-bot-api/internal/migrations"
	"github.com/brotherhood228/dating-bot-api/pkg/mongo"
	"github.com/brotherhood228/dating-bot-api/pkg/postgree"
	"github.com/brotherhood228/dating-bot-api/pkg/router"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

func main() {
	config := ConfigureAPI()
	log.WithField("CONFIG", config).Debug("App Start")

	server := router.InitRouters(routes)
	log.Info("Init routers success")

	//migrations
	pg, err := sql.Open("postgres", config.PGURI)
	if err != nil {
		log.Errorf("Bad migrations Postgres with err: %v", err)
		ShutDown()
		panic(err)
	}
	err = goose.Up(pg, "/var")
	if err != nil {
		log.Errorf("Bad migrations Postgres with err: %v", err)
		ShutDown()
		panic(err)
	}
	pg.Close()
	log.Info("Success init migrations")

	//init pg connection
	err = postgree.Init(config.PGURI)
	if err != nil {
		log.Errorf("Bad init Postgres with err: %v", err)
		ShutDown()
		panic(err)
	}
	log.Info("Success init Postgres")

	mongo.InitMongo(config.MongoURI)
	log.Info("Success init Mongo")

	go func() {
		err := server.Start(config.PORT)
		if err != nil {
			log.Info("HTTP Server Close with err: %v", err)
		}
	}()

	go health.InitUpTime()

	go StartDebugService(":9991")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-exit
	err = server.Close()
	if err != nil {
		log.Infof("Bad Close Server: %v", err)
	}
	ShutDown()
	log.Info("Success ShutDown Server")
}

//ShutDown graceful shutdown server
//close db connection
func ShutDown() {
	postgree.CloseDB()
	mongo.ShutDown()
}
