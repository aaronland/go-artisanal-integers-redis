package main

import (
	"github.com/aaronland/go-artisanal-integers-redis/engine"
	"github.com/aaronland/go-artisanal-integers/application"
	"log"
	"os"
)

func main() {

	flags := application.NewServerApplicationFlags()

	var dsn string
	flags.StringVar(&dsn, "dsn", "redis://localhost:6379", "The data source name (dsn) for connecting to Redis.")

	application.ParseFlags(flags)

	eng, err := engine.NewRedisEngine(dsn)

	if err != nil {
		log.Fatal(err)
	}

	app, err := application.NewServerApplication(eng)

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(flags)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
