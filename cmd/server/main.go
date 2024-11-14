package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/Sayonara143/urfu-pp-hhru-backend/api"
	"github.com/Sayonara143/urfu-pp-hhru-backend/confmanager"
	"github.com/Sayonara143/urfu-pp-hhru-backend/logger"
	"github.com/Sayonara143/urfu-pp-hhru-backend/services/auth"
	"github.com/Sayonara143/urfu-pp-hhru-backend/services/hh"
	"github.com/Sayonara143/urfu-pp-hhru-backend/storages/postgres"
)

const prodName = "dl-n4-1c-services"

var version = "v0.0.0"

func main() {

	pathToConf := flag.String("config", "examples/config.yaml", "Path to configuration file")
	flag.Parse()

	var c configuration
	if err := confmanager.FromYAML(&c, *pathToConf); err != nil {
		panic(err)
	}

	log := logger.New(c.App.Environment == prodName)
	log.AppendFields(map[string]interface{}{
		"version":     version,
		"environment": c.App.Environment,
	})

	// STORAGE
	db, err := postgres.New(
		c.Storages.Postgres.Addr,
		c.Storages.Postgres.Name,
		c.Storages.Postgres.User,
		c.Storages.Postgres.Pass,
		c.Storages.Postgres.Pool,
		c.App.Environment,
	)
	if err != nil {
		log.Fatal(err, "init db")
	}

	if err := db.Migrate(); err != nil {
		log.Fatal(err, "migrating db")
	}

	authSvc := auth.New(db)
	giftSvc := hh.New(db)

	// HTTP SERVER
	httpServer := api.New(giftSvc, authSvc, log)

	go func() {
		if err := httpServer.Start(c.App.HTTPAddr); err != nil {
			log.Error(err, "serving request")
		}
	}()

	log.Info("service started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	if err := db.Close(); err != nil {
		log.Error(err, "close db")
	}

	if err := httpServer.Close(context.Background()); err != nil {
		log.Error(err, "close http server")
	}

	log.Info("service shutdown")
}

type configuration struct {
	App struct {
		Environment string `yaml:"environment"`
		HTTPAddr    string `yaml:"httpAddr"`
	} `yaml:"app"`
	Storages struct {
		Postgres struct {
			Addr string `yaml:"addr"`
			Name string `yaml:"name"`
			User string `yaml:"user"`
			Pass string `yaml:"pass"`
			Pool int    `yaml:"pool"`
		} `yaml:"postgres"`
		Sentinel struct {
			Hosts      []string `yaml:"hosts"`
			Password   string   `yaml:"password"`
			DB         int      `yaml:"db"`
			Prefix     string   `yaml:"prefix"`
			IsEnabled  bool     `yaml:"isEnabled"`
			MasterName string   `yaml:"masterName"`
			PoolSize   int      `yaml:"poolSize"`
		} `yaml:"sentinel"`
		Kafka struct {
			IsEnabled bool     `yaml:"isEnabled"`
			Brokers   []string `yaml:"brokers"`
			TopicName string   `yaml:"topicName"`
			Username  string   `yaml:"username"`
			Password  string   `yaml:"password"`
		} `yaml:"kafka"`
	} `yaml:"storages"`
	Credentials map[string]string `yaml:"credentials"`
}
